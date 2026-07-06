package request_context

import (
	"github.com/funtimecoding/soil/pkg/web/constant"
	"net/http"
)

func (c *Context) WriteUnauthorized() {
	c.Write(http.StatusUnauthorized, constant.Unauthorized)
}
