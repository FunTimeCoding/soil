package site

import "github.com/funtimecoding/soil/pkg/browser/chromium/session"

func New() *Site {
	return &Site{session: session.New("chatgpt.com")}
}
