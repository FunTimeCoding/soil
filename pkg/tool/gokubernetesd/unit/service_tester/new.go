package service_tester

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/cluster"
	"k8s.io/apimachinery/pkg/runtime"
	dynamicFake "k8s.io/client-go/dynamic/fake"
	kubernetesFake "k8s.io/client-go/kubernetes/fake"
	"path/filepath"
	"testing"
)

func New(
	t *testing.T,
	o ...Option,
) *Tester {
	t.Helper()
	set := kubernetesFake.NewSimpleClientset()
	dynamic := dynamicFake.NewSimpleDynamicClientWithCustomListKinds(
		runtime.NewScheme(),
		gvrListKinds(),
	)
	c := cluster.NewWithResources("test", set, dynamic, nil, apiResources())
	result := &Tester{
		Cluster:   c,
		Clientset: set,
		Dynamic:   dynamic,
		Reporter:  memory.New(),
		clusters:  map[string]*cluster.Cluster{c.Name(): c},
		litePath:  filepath.Join(t.TempDir(), "gokubernetesd.db"),
		t:         t,
	}

	for _, f := range o {
		f(result)
	}

	return result
}
