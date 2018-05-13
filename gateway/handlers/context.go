package handlers

import (
	"github.com/KEXPCapstone/shelves-server/gateway/models/users"
	"github.com/KEXPCapstone/shelves-server/gateway/sessions"
)

type HandlerCtx struct {
	signingKey   string
	sessionStore sessions.Store
	userStore    users.UserStore
}

func NewHandlerContext(signingKey string, ss sessions.Store, us users.UserStore) *HandlerCtx {
	return &HandlerCtx{
		signingKey:   signingKey,
		sessionStore: ss,
		userStore:    us,
	}
}
