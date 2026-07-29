package request_context

import "github.com/funtimecoding/soil/pkg/web/constant"

func (c *Context) Referer() string {
	result := c.request.Referer()

	if result == "" {
		return constant.LocationRoot
	}

	return result
}
