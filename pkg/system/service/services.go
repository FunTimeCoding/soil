package service

import (
	"github.com/funtimecoding/soil/pkg/system/constant"
	"runtime"
)

func (c *Client) Services() []*Service {
	if runtime.GOOS == constant.Darwin {
		return c.launchServices()
	}

	return c.unitServices()
}
