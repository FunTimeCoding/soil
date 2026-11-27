package browser

import (
	"github.com/funtimecoding/soil/pkg/browser/chromium/protocol"
	"github.com/funtimecoding/soil/pkg/browser/face"
	"github.com/funtimecoding/soil/pkg/tool/gochromed/constant"
)

func (b *Browser) Page(identifier string) face.ChromiumProtocolClient {
	return protocol.NewIdentifier(
		b.ChromiumClient,
		identifier,
	).WithTimeout(constant.TargetTimeout)
}
