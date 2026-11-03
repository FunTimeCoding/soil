package terraform_tester

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

func build(
	t *testing.T,
	c *mock_client.Client,
	hostname string,
	unlockExit int,
	plan string,
) *Tester {
	t.Helper()
	result := fixture(t, c, hostname, unlockExit, plan, nil)
	o := result.Option
	l := logger.New(t.Context())
	m := relational.Open(l, o.PostgresLocator, o.LitePath)
	t.Cleanup(func() { relational.CloseMapper(m) })
	migrate.AutoMigrate(m, constant.RunsTable)
	s := store.New(m, constant.RunsTable)
	p := reaper.New(result.Reporter, prometheus.NewRegistry())
	p.Start()
	t.Cleanup(p.Close)
	r := runner.New(
		o,
		s,
		l,
		result.Reporter,
		p,
		prometheus.NewRegistry(),
		o.Cluster(),
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
