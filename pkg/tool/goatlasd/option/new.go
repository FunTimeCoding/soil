package option

import (
	kubernetesClient "github.com/funtimecoding/soil/pkg/kubernetes/client"
	kubernetesConstant "github.com/funtimecoding/soil/pkg/kubernetes/constant"
	kubernetes "github.com/funtimecoding/soil/pkg/kubernetes/face"
	netboxClient "github.com/funtimecoding/soil/pkg/tool/gonetboxd/client"
	netbox "github.com/funtimecoding/soil/pkg/tool/gonetboxd/face"
)

func New() *Atlas {
	return &Atlas{
		Netbox: func() netbox.Client {
			return netboxClient.NewEnvironment()
		},
		Kubernetes: func() kubernetes.Client {
			if c, e := kubernetesClient.TryInCluster(
				kubernetesConstant.InCluster,
			); e == nil {
				return c
			}

			return kubernetesClient.NewEnvironment()
		},
	}
}
