package service_tester

import (
	"github.com/funtimecoding/soil/pkg/kubernetes/client/mock_client"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/provision/migrate"
	"github.com/funtimecoding/soil/pkg/provision/store"
	"github.com/funtimecoding/soil/pkg/relational"
	"github.com/funtimecoding/soil/pkg/system/reaper"
	"github.com/funtimecoding/soil/pkg/tool/goterraformd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goterraformd/runner"
	"github.com/funtimecoding/soil/pkg/tool/goterraformd/service"
	"github.com/prometheus/client_golang/prometheus"
	"testing"
)

func New(
	t *testing.T,
	c *mock_client.Client,
	o ...Option,
) *Tester {
	t.Helper()
	result := NewFixture(t, c, o...)
	p := result.Option
	l := logger.New(t.Context())
	m := relational.Open(l, p.PostgresLocator, p.LitePath)
	t.Cleanup(func() { relational.CloseMapper(m) })
	migrate.AutoMigrate(m, constant.RunsTable)
	s := store.New(m, constant.RunsTable)
	k := reaper.New(result.Reporter, prometheus.NewRegistry())
	k.Start()
	t.Cleanup(k.Close)
	r := runner.New(
		p,
		s,
		l,
		result.Reporter,
		k,
		prometheus.NewRegistry(),
		p.Cluster(),
		nil,
	)
	r.Start()
	t.Cleanup(
		func() {
			r.Close()
			<-r.Finished()
		},
	)
	result.Service = service.New(s, r)
	result.waitForApply()

	return result
}
