package face

import "github.com/funtimecoding/soil/pkg/chromium/face"

type ChromiumSource interface {
	face.ChromiumClient
	Page(identifier string) face.ProtocolClient
}
