package handlers

import (
	"time"

	"github.com/KEXPCapstone/shelves-server/gateway/models/users"
)

type SessionState struct {
	AuthUsr   *users.User
	StartTime time.Time
}
