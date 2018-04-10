package handlers

import "github.com/KEXPCapstone/shelves-server/library/models/releases"

type HandlerCtx struct {
	releaseStore releases.ReleaseStore
}

func NewHandlerContext(rs releases.ReleaseStore) *HandlerCtx {
	return &HandlerCtx{
		releaseStore: rs,
	}
}
