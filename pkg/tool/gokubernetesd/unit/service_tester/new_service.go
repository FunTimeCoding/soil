package service_tester

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/relational"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/option"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/port_forward/registry"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/service"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/store"
	"testing"
)

func NewService(
	t *testing.T,
	o *option.Server,
) *service.Service {
	t.Helper()
	l := logger.New(t.Context())
	r := memory.New()
	d := relational.Open(l, o.PostgresLocator, o.LitePath)
	t.Cleanup(func() { relational.CloseMapper(d) })
	clusters := o.Clusters()
	guards := map[string]*peer.Guard{}

	for name, c := range clusters {
		g := peer.New(join.Colon(constant.ClusterPeer, name), c.Host(), c, l, r)
		t.Cleanup(g.Close)
		guards[name] = g
	}

	forwards := registry.New()
	t.Cleanup(forwards.Close)

	return service.New(clusters, guards, store.New(d), forwards, r)
}
