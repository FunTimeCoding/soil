package site

import (
	"github.com/funtimecoding/soil/pkg/chromium/session"
	"github.com/funtimecoding/soil/pkg/generative/anthropic/site/reader"
)

type Site struct {
	*reader.Reader
	session *session.Session
}
