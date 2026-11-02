package service_tester

import (
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/cluster"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/option"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/service"
	dynamicFake "k8s.io/client-go/dynamic/fake"
	kubernetesFake "k8s.io/client-go/kubernetes/fake"
	"testing"
)

type Tester struct {
	Option    *option.Server
	Service   *service.Service
	Cluster   *cluster.Cluster
	Clientset *kubernetesFake.Clientset
	Dynamic   *dynamicFake.FakeDynamicClient
	t         *testing.T
}
