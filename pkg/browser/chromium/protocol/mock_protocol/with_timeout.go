package mock_protocol

import (
	"github.com/funtimecoding/soil/pkg/browser/face"
	"time"
)

func (p *Protocol) WithTimeout(_ time.Duration) face.ChromiumProtocolClient {
	return p
}
