package ansible

import (
	"github.com/funtimecoding/soil/pkg/provision/ansible/constant"
	"github.com/funtimecoding/soil/pkg/system/environment"
)

func NewEnvironment() *Client {
	return New(environment.Required(constant.InventoryEnvironment))
}
