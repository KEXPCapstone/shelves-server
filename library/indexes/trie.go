package indexes

import (
	"errors"
	"sort"
	"sync"
	"unicode/utf8"

	"gopkg.in/mgo.v2/bson"
)

type TrieNode struct {
	Key           rune
	Children      map[rune]*TrieNode
	SearchResults []SearchResult
	Parent        *TrieNode
	mx            sync.RWMutex
}

type SearchResult struct {
	ReleaseID      bson.ObjectId `json:"releaseID"`
	FieldMatchedOn string        `json:"fieldMatchedOn"`
}

func CreateTrieRoot() *TrieNode {
	return &TrieNode{Children: make(map[rune]*TrieNode)}
}

func (t *TrieNode) AddToTrie(prefix string, searchResult SearchResult) {

	t.mx.Lock()
	curr := t
	for _, char := range prefix {
		if curr.Children == nil {
			curr.Children = make(map[rune]*TrieNode)
		}

		child, ok := curr.Children[char]
		if !ok {
			child = &TrieNode{Key: char, Parent: curr}
			curr.Children[char] = child // add new node to trie if this part of prefix not in trie
		}
		curr = child
	}
	if !t.nodeContainsRelease(curr, searchResult.ReleaseID) {
		curr.SearchResults = append(curr.SearchResults, searchResult)
	}
	t.mx.Unlock()

}

func (t *TrieNode) nodeContainsRelease(node *TrieNode, releaseID bson.ObjectId) bool {
	for _, val := range node.SearchResults {
		if val.ReleaseID == releaseID {
			return true
		}
	}
	return false
}

func (t *TrieNode) SearchReleases(prefix string, maxResults int) []SearchResult {
	t.mx.RLock()
	defer t.mx.RUnlock()
	startNode, err := t.searchPrefixForNode(prefix)
	if err != nil {
		return []SearchResult{}
	}
	return t.findResultsFromPrefix(startNode, maxResults, []SearchResult{}, make(map[bson.ObjectId]bool))
}

func (t *TrieNode) searchPrefixForNode(prefix string) (*TrieNode, error) {
	curr := t
	for _, char := range prefix {
		child, ok := curr.Children[char]

		if ok {
			curr = child

		} else {
			return nil, errors.New("Prefix doesn't exist in trie")
		}
	}
	return curr, nil
}

func (t *TrieNode) findResultsFromPrefix(node *TrieNode, maxResults int, results []SearchResult, idsInResults map[bson.ObjectId]bool) []SearchResult {
	for _, searchResult := range node.SearchResults {

		_, ok := idsInResults[searchResult.ReleaseID]
		if len(results) < maxResults && !ok {
			idsInResults[searchResult.ReleaseID] = true
			results = append(results, searchResult)
		} else {
			return results
		}
	}
	// sort the keys of the map so we iterate consistently
	var sortedKeys []string
	for key := range node.Children {
		sortedKeys = append(sortedKeys, string(key))
	}
	sort.Strings(sortedKeys)
	for _, key := range sortedKeys {
		r, _ := utf8.DecodeRuneInString(key)
		childNode := node.Children[r]
		results = t.findResultsFromPrefix(childNode, maxResults, results, idsInResults)
	}
	return results
}

func (t *TrieNode) RemoveKeyAndReleaseID(key string, value bson.ObjectId) error {
	t.mx.Lock()
	defer t.mx.Unlock()
	node, err := t.searchPrefixForNode(key)
	if err != nil { // case: key is not in trie
		return err
	}
	found := false
	for i, searchResult := range node.SearchResults {
		if searchResult.ReleaseID == value {
			node.SearchResults = append(node.SearchResults[:i], node.SearchResults[i+1:]...)
			found = true
		}
	}
	if !found { // very rarely thrown -- would indicate either a) user wasn't added to trie correctly, or b) did not correctly clean up the trie on key/val deletion
		return errors.New("No release exists at the provided key")
	}

	curr := node
	for curr.Key != 0 && len(curr.SearchResults) == 0 && len(curr.Children) == 0 { // check if curr.Key is empty (for the root node)
		// remove current from the parent's list of children
		delete(curr.Parent.Children, curr.Key)
		curr = curr.Parent
	}
	return nil
}
