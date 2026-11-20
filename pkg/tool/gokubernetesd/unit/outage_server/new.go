package outage_server

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/cluster"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/unit/base"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/unit/service_tester"
	"github.com/funtimecoding/soil/pkg/web/outage_tester"
	"testing"
)

func New(
	t *testing.T,
	r *memory.Memory,
) (*model_context_server.Server, func() error) {
	t.Helper()
	c := service_tester.RefusingCluster("test")
	s := base.New(
		t,
		service_tester.WithClusters(map[string]*cluster.Cluster{c.Name(): c}),
		service_tester.WithArmed(),
		service_tester.WithReporter(r),
	)

	return s.Served(), outage_tester.New(t).Guard(
		s.Daemon().Workers(),
		join.Colon(constant.ClusterPeer, c.Name()),
	).Unavailable
}
