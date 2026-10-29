package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/service/cluster"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/service/request"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/unit/recovering_clientset"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/unit/service_tester"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"testing"
)

func TestAnUnreachableClusterFailsFastWithoutBeingCalledAgain(t *testing.T) {
	v := service_tester.New(t)
	v.RefuseAll(refusedConnection())
	list := request.List{ResourceType: "pod", Namespace: "default"}
	_, e := v.Service.ListResources(t.Context(), "test", list)
	assert.True(t, unavailable.Is(e))
	calls := v.DynamicCalls()
	_, f := v.Service.ListResources(t.Context(), "test", list)
	assert.True(t, unavailable.Is(f))
	assert.StringContains(t, "kubernetes:test unreachable since", f.Error())
	assert.Integer(t, calls, v.DynamicCalls())
}

func TestOneUnreachableClusterLeavesTheOthersAnswering(t *testing.T) {
	down := service_tester.New(t)
	down.RefuseAll(refusedConnection())
	up := service_tester.New(t)
	s := service_tester.NewService(
		t,
		map[string]*cluster.Cluster{"test": down.Cluster, "lima": up.Cluster},
	)
	list := request.List{ResourceType: "pod", Namespace: "default"}
	_, e := s.ListResources(t.Context(), "test", list)
	assert.True(t, unavailable.Is(e))
	_, f := s.ListResources(t.Context(), "lima", list)
	assert.FatalOnError(t, f)
}

func TestAClusterThatFailsDiscoveryRecoversOnTheNextCall(t *testing.T) {
	c := cluster.New(
		"mike",
		recovering_clientset.New(
			1,
			[]*v1.APIResourceList{
				{
					GroupVersion: "v1",
					APIResources: []v1.APIResource{
						{Name: "pods", Namespaced: true, Kind: "Pod"},
					},
				},
			},
		),
		nil,
		nil,
	)
	_, _, e := c.ResolveResource("pod")
	assert.NotNil(t, e)
	_, namespaced, f := c.ResolveResource("pod")
	assert.FatalOnError(t, f)
	assert.True(t, namespaced)
}

func TestTheServiceHandsOneGuardPerClusterToTheLifecycle(t *testing.T) {
	a := service_tester.New(t)
	s := service_tester.NewService(
		t,
		map[string]*cluster.Cluster{"test": a.Cluster, "november": a.Cluster},
	)
	assert.Count(t, 2, s.Workers())
}
