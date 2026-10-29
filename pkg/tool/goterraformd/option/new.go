package option

import (
	"github.com/funtimecoding/soil/pkg/kubernetes/client"
	kubernetesConstant "github.com/funtimecoding/soil/pkg/kubernetes/constant"
	"github.com/funtimecoding/soil/pkg/kubernetes/face"
	"github.com/funtimecoding/soil/pkg/tool/goterraformd/constant"
)

func New() *Terraform {
	return &Terraform{
		Executable: constant.Command,
		Cluster: func() face.Client {
			return client.NewInCluster(kubernetesConstant.InCluster)
		},
	}
}
