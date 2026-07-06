package request_context

import "github.com/funtimecoding/soil/pkg/web/constant"

func (c *Context) SetLastLocation() {
	// Ignore if the referer header is set, favicon.ico is a common example.
	if _, okay := c.Header()[constant.Referer]; okay {
		return
	}

	c.SetCookie(constant.LastLocation, c.request.URL.Path)
}
