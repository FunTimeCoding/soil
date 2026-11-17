package service_tester

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/cluster"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/daemon"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/service"
	dynamicFake "k8s.io/client-go/dynamic/fake"
	kubernetesFake "k8s.io/client-go/kubernetes/fake"
	"testing"
)

type Tester struct {
	Cluster   *cluster.Cluster
	Clientset *kubernetesFake.Clientset
	Dynamic   *dynamicFake.FakeDynamicClient
	Reporter  *memory.Memory
	clusters  map[string]*cluster.Cluster
	arm       bool
	litePath  string
	service   *service.Service
	daemon    *daemon.Daemon
	t         *testing.T
}
