package face

import "github.com/funtimecoding/soil/pkg/browser/face"

type ChromiumSource interface {
	face.ChromiumClient
	Page(identifier string) face.ChromiumProtocolClient
}
