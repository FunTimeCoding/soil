package option

import (
	"github.com/funtimecoding/soil/pkg/kubernetes/client"
	kubernetesConstant "github.com/funtimecoding/soil/pkg/kubernetes/constant"
	"github.com/funtimecoding/soil/pkg/kubernetes/face"
	"github.com/funtimecoding/soil/pkg/network"
	provisionClient "github.com/funtimecoding/soil/pkg/provision/client"
	provisionFace "github.com/funtimecoding/soil/pkg/provision/face"
	"github.com/funtimecoding/soil/pkg/tool/goterraformd/constant"
)

func New() *Terraform {
	result := &Terraform{
		Executable: constant.Command,
		Cluster: func() face.Client {
			return client.NewInCluster(kubernetesConstant.InCluster)
		},
	}
	result.DownstreamClients = func() []provisionFace.Client {
		var clients []provisionFace.Client

		for _, address := range result.Downstream {
			host, port := network.SplitHostPort(address)
			clients = append(
				clients,
				provisionClient.New(host, port, result.DownstreamToken),
			)
		}

		return clients
	}

	return result
}
