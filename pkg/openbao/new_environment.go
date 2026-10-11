package openbao

import (
	"github.com/funtimecoding/soil/pkg/openbao/constant"
	"github.com/funtimecoding/soil/pkg/system/environment"
)

func NewEnvironment() *Client {
	return New(environment.Required(constant.HostEnvironment))
}
