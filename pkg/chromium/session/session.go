package session

import (
	"github.com/funtimecoding/soil/pkg/chromium"
	"github.com/funtimecoding/soil/pkg/chromium/protocol"
)

type Session struct {
	*protocol.Protocol
	client *chromium.Client
}
