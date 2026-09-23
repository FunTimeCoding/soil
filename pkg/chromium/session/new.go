package session

import (
	"github.com/funtimecoding/soil/pkg/chromium"
	"github.com/funtimecoding/soil/pkg/chromium/constant"
	"github.com/funtimecoding/soil/pkg/chromium/protocol"
)

func New(tab string) *Session {
	c := chromium.NewEnvironment()

	if c.TabByHost(tab) == nil {
		c.Close()
		panic(constant.TabNotFound)
	}

	return &Session{Protocol: protocol.New(c, tab), client: c}
}
