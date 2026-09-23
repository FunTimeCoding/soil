package browser

import "github.com/funtimecoding/soil/pkg/chromium"

func New(c *chromium.Client) *Browser {
	return &Browser{Client: c}
}
