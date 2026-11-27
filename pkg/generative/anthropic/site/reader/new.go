package reader

import "github.com/funtimecoding/soil/pkg/browser/chromium/protocol"

func New(p *protocol.Protocol) *Reader {
	return &Reader{Protocol: p}
}
