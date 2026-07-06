package service_tester

import (
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/integration_test/store_tester"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/service"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/service/cluster"
	dynamicFake "k8s.io/client-go/dynamic/fake"
	kubernetesFake "k8s.io/client-go/kubernetes/fake"
)

type Tester struct {
	*store_tester.Tester
	Service   *service.Service
	Cluster   *cluster.Cluster
	Clientset *kubernetesFake.Clientset
	Dynamic   *dynamicFake.FakeDynamicClient
}
