package mock_protocol

import (
	"github.com/funtimecoding/soil/pkg/chromium/face"
	"time"
)

func (p *Protocol) WithTimeout(_ time.Duration) face.ProtocolClient {
	return p
}
