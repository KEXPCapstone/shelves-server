package handlers

import "github.com/KEXPCapstone/shelves-server/gateway/sessions"

type HandlerCtx struct {
	signingKey   string
	sessionStore sessions.Store
	// TODO
}

func NewHandlerContext(signingKey string, ss sessions.Store) *HandlerCtx {
	// TODO
	return &HandlerCtx{
		signingKey:   signingKey,
		sessionStore: ss,
	}
}
