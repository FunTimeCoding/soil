package protocol

import "github.com/funtimecoding/soil/pkg/chromium/face"

func NewIdentifier(
	c face.ChromiumClient,
	identifier string,
) *Protocol {
	return &Protocol{client: c, identifier: identifier}
}
