package service_tester

import (
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/service"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/service/cluster"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/unit/store_tester"
	dynamicFake "k8s.io/client-go/dynamic/fake"
	kubernetesFake "k8s.io/client-go/kubernetes/fake"
	"testing"
)

type Tester struct {
	*store_tester.Tester
	Service   *service.Service
	Cluster   *cluster.Cluster
	Clientset *kubernetesFake.Clientset
	Dynamic   *dynamicFake.FakeDynamicClient
	t         *testing.T
}
