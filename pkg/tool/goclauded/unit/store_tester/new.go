package store_tester

import (
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/relational"
	"github.com/funtimecoding/soil/pkg/relational/lite"
	"github.com/funtimecoding/soil/pkg/time/clock_tester"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/migrate"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store"
	"testing"
	"time"
)

func New(
	t *testing.T,
	o ...Option,
) *Tester {
	t.Helper()
	s := &setup{zone: time.UTC}

	for _, f := range o {
		f(s)
	}

	c := clock_tester.NewAt(time.Now().In(s.zone))
	d := lite.NewMemory()

	if s.path != "" {
		d = lite.New(logger.New(t.Context()), s.path)
	}

	t.Cleanup(func() { relational.CloseMapper(d) })
	migrate.AutoMigrate(d)

	return &Tester{t: t, mapper: d, Store: store.New(d, c.Now), clock: c}
}
