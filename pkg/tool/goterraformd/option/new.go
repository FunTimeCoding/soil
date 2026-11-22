package option

import (
	"github.com/funtimecoding/soil/pkg/kubernetes/client"
	kubernetesConstant "github.com/funtimecoding/soil/pkg/kubernetes/constant"
	"github.com/funtimecoding/soil/pkg/kubernetes/face"
	"github.com/funtimecoding/soil/pkg/network"
	provisionClient "github.com/funtimecoding/soil/pkg/provision/client"
	provisionFace "github.com/funtimecoding/soil/pkg/provision/face"
	"github.com/funtimecoding/soil/pkg/tool/goterraformd/constant"
	"github.com/funtimecoding/soil/pkg/web/locator"
)

func New() *Terraform {
	result := &Terraform{
		Executable: constant.Command,
		Cluster: func() face.KubernetesClient {
			return client.NewInCluster(kubernetesConstant.InCluster)
		},
	}
	result.DownstreamClients = func() []provisionFace.Client {
		var clients []provisionFace.Client

		for _, address := range result.Downstream {
			host, port := network.SplitHostPort(address)
			clients = append(
				clients,
				provisionClient.New(
					locator.New(host).Port(port).Insecure(),
					result.DownstreamToken,
				))
		}

		return clients
	}

	return result
}
