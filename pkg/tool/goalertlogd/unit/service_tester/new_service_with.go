package service_tester

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	soilFace "github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/prometheus/face"
	"github.com/funtimecoding/soil/pkg/relational"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/service"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/store"
	"testing"
	"time"
)

func NewServiceWith(
	t *testing.T,
	path string,
	a face.AlertmanagerClient,
	retention time.Duration,
	clock func() time.Time,
	n soilFace.EventNotifier,
) *service.Service {
	t.Helper()
	l := logger.New(t.Context())
	d := relational.Open(l, "", path)
	t.Cleanup(func() { relational.CloseMapper(d) })
	g := peer.New(constant.AlertmanagerPeer, a.Address(), a, l, memory.New())
	t.Cleanup(g.Close)

	return service.New(store.New(d), a, g, retention, clock, n)
}
