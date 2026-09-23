package browser

import (
	"github.com/funtimecoding/soil/pkg/chromium/protocol"
	"github.com/funtimecoding/soil/pkg/tool/gochromed/face"
)

func (b *Browser) Page(identifier string) face.Page {
	return protocol.NewIdentifier(b.Client, identifier)
}
