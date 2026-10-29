package service_tester

import (
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/service/cluster"
	"k8s.io/apimachinery/pkg/runtime"
	dynamicFake "k8s.io/client-go/dynamic/fake"
	kubernetesFake "k8s.io/client-go/kubernetes/fake"
	"testing"
)

func New(t *testing.T) *Tester {
	t.Helper()
	set := kubernetesFake.NewSimpleClientset()
	dynamic := dynamicFake.NewSimpleDynamicClientWithCustomListKinds(
		runtime.NewScheme(),
		gvrListKinds(),
	)
	c := cluster.NewWithResources("test", set, dynamic, nil, apiResources())

	return &Tester{
		t:         t,
		Service:   NewService(t, map[string]*cluster.Cluster{c.Name(): c}),
		Cluster:   c,
		Clientset: set,
		Dynamic:   dynamic,
	}
}
