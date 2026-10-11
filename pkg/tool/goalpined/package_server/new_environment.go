package package_server

import (
	"github.com/funtimecoding/soil/pkg/system/environment"
	"github.com/funtimecoding/soil/pkg/tool/goalpined/constant"
)

func NewEnvironment() *Server {
	return New(environment.Required(constant.SignatureKeyEnvironment))
}
