package service_tester

import (
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/relational"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/port_forward/registry"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/service"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/store"
)

func (t *Tester) Service() *service.Service {
	t.t.Helper()

	if t.service != nil {
		return t.service
	}

	l := logger.New(t.t.Context())
	d := relational.Open(l, "", t.litePath)
	t.t.Cleanup(func() { relational.CloseMapper(d) })
	guards := map[string]*peer.Guard{}

	for name, c := range t.clusters {
		g := peer.New(
			join.Colon(constant.ClusterPeer, name),
			c.Host(),
			c,
			l,
			t.Reporter,
		)
		t.t.Cleanup(g.Close)
		guards[name] = g
	}

	forwards := registry.New()
	t.t.Cleanup(forwards.Close)
	t.service = service.New(
		t.clusters,
		guards,
		store.New(d),
		forwards,
		t.Reporter,
	)

	return t.service
}
