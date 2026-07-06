package request_context

import "github.com/funtimecoding/soil/pkg/web/constant"

func (c *Context) WriteSuccess() {
	c.WriteOkay(constant.Success)
}
