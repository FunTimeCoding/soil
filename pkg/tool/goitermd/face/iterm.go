package face

import (
	"github.com/funtimecoding/soil/pkg/iterm/history"
	"github.com/funtimecoding/soil/pkg/iterm/screen"
	"github.com/funtimecoding/soil/pkg/iterm/session"
	"time"
)

type Iterm interface {
	Sessions() ([]*session.Session, error)
	ReadScreen(identifier string) (*screen.Screen, error)
	ReadHistory(
		identifier string,
		count int,
	) (*history.History, error)
	SendText(
		identifier string,
		text string,
		enter bool,
	) error
	SendKeys(
		identifier string,
		keys []string,
		interval time.Duration,
	) error
	SetTabTitle(
		tabIdentifier string,
		title string,
	) error
	SetTabColor(
		sessionIdentifier string,
		red int,
		green int,
		blue int,
	) error
	CreateTab() (*session.Session, error)
}
