package option

import (
	kubernetesClient "github.com/funtimecoding/soil/pkg/kubernetes/client"
	kubernetesConstant "github.com/funtimecoding/soil/pkg/kubernetes/constant"
	kubernetes "github.com/funtimecoding/soil/pkg/kubernetes/face"
	netboxClient "github.com/funtimecoding/soil/pkg/tool/gonetboxd/client"
	netbox "github.com/funtimecoding/soil/pkg/tool/gonetboxd/face"
	authorization "github.com/funtimecoding/soil/pkg/web/authorization/client"
	"github.com/funtimecoding/soil/pkg/web/option"
)

func New() *Atlas {
	return &Atlas{
		Web: option.Web{Authorization: authorization.NewEnvironment},
		Netbox: func() netbox.Client {
			return netboxClient.NewEnvironment()
		},
		Kubernetes: func() kubernetes.KubernetesClient {
			if c, e := kubernetesClient.TryInCluster(
				kubernetesConstant.InCluster,
			); e == nil {
				return c
			}

			return kubernetesClient.NewEnvironment()
		},
	}
}
