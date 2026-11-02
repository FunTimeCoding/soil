package option

import (
	"github.com/funtimecoding/soil/pkg/provision/salt"
	"github.com/funtimecoding/soil/pkg/tool/gosaltd/constant"
)

func New() *Salt {
	return &Salt{
		KeySyncInterval: constant.KeySyncInterval,
		Master:          salt.NewEnvironment,
	}
}
