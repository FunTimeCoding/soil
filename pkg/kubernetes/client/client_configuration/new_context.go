package client_configuration

import (
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/system/constant"
	"github.com/funtimecoding/soil/pkg/system/join"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func NewContext(context string) (*rest.Config, error) {
	return clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
		&clientcmd.ClientConfigLoadingRules{
			ExplicitPath: join.Absolute(
				system.Home(),
				constant.KubernetesConfigurationPath,
			),
		},
		&clientcmd.ConfigOverrides{CurrentContext: context},
	).ClientConfig()
}
