package handlers

import (
	"github.com/KEXPCapstone/shelves-server/library/indexes"
	"github.com/KEXPCapstone/shelves-server/library/models/releases"
)

type HandlerCtx struct {
	libraryStore releases.LibraryStore
	releaseTrie  *indexes.TrieNode
}

func NewHandlerContext(ls releases.LibraryStore, rt *indexes.TrieNode) *HandlerCtx {
	return &HandlerCtx{
		libraryStore: ls,
		releaseTrie:  rt,
	}
}
