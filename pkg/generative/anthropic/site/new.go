package site

import "github.com/funtimecoding/soil/pkg/chromium/protocol"

func New() *Site {
	return &Site{protocol: protocol.New("#settings/usage")}
}
