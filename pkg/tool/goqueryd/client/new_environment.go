package client

import (
	"github.com/funtimecoding/soil/pkg/system/environment"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"
)

func NewEnvironment() *Client {
	return New(
		environmentLocator(),
		environment.Required(constant.TokenEnvironment),
	)
}
