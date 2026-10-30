package client

import (
	"github.com/funtimecoding/soil/pkg/system/environment"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"
)

func NewLongEnvironment() *Client {
	return NewLong(
		environmentLocator(),
		environment.Required(constant.TokenEnvironment),
	)
}
