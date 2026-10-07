package service

import (
	"github.com/funtimecoding/soil/pkg/system/constant"
	"github.com/funtimecoding/soil/pkg/system/types/service"
	"runtime"
)

func (c *Client) Services() []*service.Service {
	if runtime.GOOS == constant.Darwin {
		return c.launchServices()
	}

	return c.unitServices()
}
