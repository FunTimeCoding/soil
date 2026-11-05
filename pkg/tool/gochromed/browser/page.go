package browser

import (
	"github.com/funtimecoding/soil/pkg/chromium/face"
	"github.com/funtimecoding/soil/pkg/chromium/protocol"
	"github.com/funtimecoding/soil/pkg/tool/gochromed/constant"
)

func (b *Browser) Page(identifier string) face.ProtocolClient {
	return protocol.NewIdentifier(
		b.ChromiumClient,
		identifier,
	).WithTimeout(constant.TargetTimeout)
}
