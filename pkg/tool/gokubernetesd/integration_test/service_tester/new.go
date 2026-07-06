package service_tester

import (
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/integration_test/store_tester"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/service"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/service/cluster"
	"k8s.io/apimachinery/pkg/runtime"
	dynamicFake "k8s.io/client-go/dynamic/fake"
	kubernetesFake "k8s.io/client-go/kubernetes/fake"
	"testing"
)

func New(t *testing.T) *Tester {
	t.Helper()
	s := store_tester.New(t)
	scheme := runtime.NewScheme()
	clientset := kubernetesFake.NewSimpleClientset()
	dynamic := dynamicFake.NewSimpleDynamicClientWithCustomListKinds(
		scheme,
		gvrListKinds(),
	)
	c := cluster.NewWithResources(
		"test",
		clientset,
		dynamic,
		nil,
		apiResources(),
	)
	svc := service.NewWithCluster(s.Store, c)

	return &Tester{
		Tester:    s,
		Service:   svc,
		Cluster:   c,
		Clientset: clientset,
		Dynamic:   dynamic,
	}
}
