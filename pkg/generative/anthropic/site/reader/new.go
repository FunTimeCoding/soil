package reader

import "github.com/funtimecoding/soil/pkg/chromium/protocol"

func New(p *protocol.Protocol) *Reader {
	return &Reader{Protocol: p}
}
