package protocol

import (
	"github.com/funtimecoding/soil/pkg/chromium/face"
	"time"
)

func (p *Protocol) WithTimeout(d time.Duration) face.ProtocolClient {
	p.timeout = d

	return p
}
