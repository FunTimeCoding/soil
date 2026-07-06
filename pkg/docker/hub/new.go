package hub

import (
	"github.com/funtimecoding/soil/pkg/docker/hub/constant"
	"github.com/funtimecoding/soil/pkg/web/locator"
)

func New() *Client {
	return &Client{
		base: locator.New(constant.Host).Base(constant.BasePath),
	}
}
