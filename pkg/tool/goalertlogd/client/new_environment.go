package client

import (
	"github.com/funtimecoding/soil/pkg/system/environment"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/constant"
)

func NewEnvironment() *Client {
	return New(environment.Required(constant.HostEnvironment))
}
