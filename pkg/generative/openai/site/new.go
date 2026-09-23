package site

import "github.com/funtimecoding/soil/pkg/chromium/session"

func New() *Site {
	return &Site{session: session.New("chatgpt.com")}
}
