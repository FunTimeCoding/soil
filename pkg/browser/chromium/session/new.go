package session

import (
	"github.com/funtimecoding/soil/pkg/browser/chromium"
	"github.com/funtimecoding/soil/pkg/browser/chromium/protocol"
	"github.com/funtimecoding/soil/pkg/browser/constant"
)

func New(tab string) *Session {
	c := chromium.NewEnvironment()

	if c.MustHostTab(tab) == nil {
		c.Close()
		panic(constant.ChromiumTabNotFound)
	}

	return &Session{Protocol: protocol.New(c, tab), client: c}
}
