package handlers

import "github.com/KEXPCapstone/shelves-server/shelves/models"

type HandlerCtx struct {
	shelfStore models.ShelfStore
}

func NewHandlerContext(store models.ShelfStore) *HandlerCtx {
	return &HandlerCtx{
		shelfStore: store,
	}
}
