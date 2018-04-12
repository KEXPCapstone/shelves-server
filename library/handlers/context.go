package handlers

import (
	"github.com/KEXPCapstone/shelves-server/library/indexes"
	"github.com/KEXPCapstone/shelves-server/library/models/releases"
)

type HandlerCtx struct {
	releaseStore releases.ReleaseStore
	releaseTrie  *indexes.TrieNode
}

func NewHandlerContext(rs releases.ReleaseStore, rt *indexes.TrieNode) *HandlerCtx {
	return &HandlerCtx{
		releaseStore: rs,
		releaseTrie:  rt,
	}
}
