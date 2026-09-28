package kubernetes

import (
	"github.com/funtimecoding/soil/pkg/kubernetes/client"
	"github.com/funtimecoding/soil/pkg/kubernetes/constant"
)

func NewEnvironment() *Collector {
	if c, e := client.TryInCluster(constant.InCluster); e == nil {
		return New(c)
	}

	return New(client.NewEnvironment())
}
