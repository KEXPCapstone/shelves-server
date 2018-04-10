package main

type HandlerCtx struct {
	releaseStore ReleaseStore
}

func NewHandlerContext(rs ReleaseStore) *HandlerCtx {
	return &HandlerCtx{
		releaseStore: rs,
	}
}
