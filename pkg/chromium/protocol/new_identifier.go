package protocol

import "github.com/funtimecoding/soil/pkg/chromium"

func NewIdentifier(
	c *chromium.Client,
	identifier string,
) *Protocol {
	return &Protocol{client: c, identifier: identifier}
}
