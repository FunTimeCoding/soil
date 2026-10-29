package face

import (
	"github.com/funtimecoding/soil/pkg/provision/face"
	"github.com/funtimecoding/soil/pkg/provision/salt/basic/response"
)

type Salt interface {
	face.Provision
	Keys() (*response.KeysReturn, error)
	Minions() ([]response.Minion, error)
}
