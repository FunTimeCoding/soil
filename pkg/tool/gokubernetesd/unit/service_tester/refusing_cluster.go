package service_tester

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/cluster"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
)

func RefusingCluster(name string) *cluster.Cluster {
	c := refusingConfiguration()
	set, e := kubernetes.NewForConfig(c)
	errors.PanicOnError(e)
	d, f := dynamic.NewForConfig(c)
	errors.PanicOnError(f)

	return cluster.NewWithResources(name, set, d, c, apiResources())
}
