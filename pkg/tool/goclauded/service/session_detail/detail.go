package session_detail

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/completion"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/session"
)

type Detail struct {
	Identifier  string
	Slug        string
	Created     string
	Alias       string
	Description string
	TurnCount   int
	Completions []completion.Completion
	Summary     string
	Session     *session.Session
}
