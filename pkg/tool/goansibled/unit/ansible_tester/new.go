package ansible_tester

import (
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/provision/migrate"
	"github.com/funtimecoding/soil/pkg/provision/store"
	"github.com/funtimecoding/soil/pkg/relational"
	"github.com/funtimecoding/soil/pkg/system/reaper"
	"github.com/funtimecoding/soil/pkg/tool/goansibled/constant"
	"github.com/funtimecoding/soil/pkg/tool/goansibled/runner"
	"github.com/funtimecoding/soil/pkg/tool/goansibled/service"
	"net/http"
	"testing"
)

func New(
	t *testing.T,
	inventory http.Handler,
) *Tester {
	t.Helper()
	result := NewFixture(t, inventory)
	o := result.Option
	l := logger.New(t.Context())
	m := relational.Open(l, o.PostgresLocator, o.LitePath)
	t.Cleanup(func() { relational.CloseMapper(m) })
	migrate.AutoMigrate(m, constant.RunsTable)
	s := store.New(m, constant.RunsTable)
	p := reaper.New(result.Reporter)
	p.Start()
	t.Cleanup(p.Close)
	r := runner.New(o, s, l, result.Reporter, p)

	if o.Inventory != nil {
		c := o.Inventory()
		t.Cleanup(c.Close)
		r.WithInventory(c)
	}

	r.Start()
	t.Cleanup(
		func() {
			r.Close()
			<-r.Finished()
		},
	)
	result.Service = service.New(s, r)

	return result
}
