package face

import (
	"github.com/funtimecoding/soil/pkg/iterm/history"
	"github.com/funtimecoding/soil/pkg/iterm/screen"
	"github.com/funtimecoding/soil/pkg/iterm/session"
)

type ItermSource interface {
	Sessions() ([]*session.Session, error)
	CreateTab() (*session.Session, error)
	ReadScreen(identifier string) (*screen.Screen, error)
	ReadHistory(
		identifier string,
		count int,
	) (*history.History, error)
	SendText(
		identifier string,
		text string,
	) error
	SendKey(
		identifier string,
		key string,
	) error
	SetTabColor(
		sessionIdentifier string,
		red int,
		green int,
		blue int,
	) error
	SetTabTitle(
		tabIdentifier string,
		title string,
	) error
	MustSessions() []*session.Session
	MustCreateTab() *session.Session
	MustReadScreen(identifier string) *screen.Screen
	MustReadHistory(
		identifier string,
		count int,
	) *history.History
}
