package package_server

import (
	"github.com/funtimecoding/soil/pkg/alpine/constant"
	"github.com/funtimecoding/soil/pkg/system/environment"
)

func NewEnvironment() *Server {
	return New(environment.Required(constant.SignatureKeyEnvironment))
}
