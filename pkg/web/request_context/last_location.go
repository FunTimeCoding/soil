package request_context

import "github.com/funtimecoding/soil/pkg/web/constant"

func (c *Context) LastLocation() string {
	if s := c.Cookie(constant.LastLocation); s != nil {
		return s.Value
	}

	return constant.LocationRoot
}
