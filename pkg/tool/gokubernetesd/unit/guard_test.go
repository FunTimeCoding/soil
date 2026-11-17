package unit

import (
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/cluster"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/unit/base"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/unit/service_tester"
	"testing"
)

func TestGuard(t *testing.T) {
	v := base.New(
		t,
		service_tester.WithClusters(map[string]*cluster.Cluster{}),
	).Served()
	v.VerifyBase(t)
	v.VerifyModelContext(t)
}
