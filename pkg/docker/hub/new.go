package hub

import (
	"github.com/funtimecoding/soil/pkg/docker/constant"
	"github.com/funtimecoding/soil/pkg/web/locator"
)

func New() *Client {
	return NewWithBase(locator.New(constant.Host).Base(constant.BasePath))
}
