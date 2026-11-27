package session

import (
	"github.com/funtimecoding/soil/pkg/browser/chromium"
	"github.com/funtimecoding/soil/pkg/browser/chromium/protocol"
)

type Session struct {
	*protocol.Protocol
	client *chromium.Client
}
