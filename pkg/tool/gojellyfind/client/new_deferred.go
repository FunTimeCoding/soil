package client

import (
	"github.com/funtimecoding/soil/pkg/tool/gojellyfind/constant"
	"github.com/funtimecoding/soil/pkg/web/locator"
)

func NewDeferred(read func(string) string) *Client {
	return newOnToken(
		locator.Environment(
			constant.HostEnvironment,
			constant.PortEnvironment,
			constant.InsecureEnvironment,
		),
		func() string { return read(constant.TokenEnvironment) },
	)
}
