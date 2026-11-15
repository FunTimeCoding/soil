package client

import (
	"github.com/funtimecoding/soil/pkg/system/environment"
	gopostgresConstant "github.com/funtimecoding/soil/pkg/tool/gopostgres/constant"
	"github.com/funtimecoding/soil/pkg/web/locator"
)

func NewEnvironment(o ...Option) *Client {
	return New(
		locator.Environment(
			gopostgresConstant.HostEnvironment,
			gopostgresConstant.PortEnvironment,
			gopostgresConstant.InsecureEnvironment,
		),
		environment.Required(gopostgresConstant.TokenEnvironment),
		o...,
	)
}
