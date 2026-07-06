package client_configuration

import (
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/system/constant"
	"github.com/funtimecoding/soil/pkg/system/join"
)

func path() string {
	return join.Absolute(system.Home(), constant.KubernetesConfigurationPath)
}
