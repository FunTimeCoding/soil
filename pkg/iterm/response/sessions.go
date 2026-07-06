package response

import "github.com/funtimecoding/soil/pkg/iterm/session"

type Sessions struct {
	Sessions []*session.Session `json:"sessions"`
}
