package protocol

import (
	"github.com/funtimecoding/soil/pkg/browser/face"
	"time"
)

func (p *Protocol) WithTimeout(d time.Duration) face.ChromiumProtocolClient {
	p.timeout = d

	return p
}
