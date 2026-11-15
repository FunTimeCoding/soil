package store_tester

import (
	"github.com/funtimecoding/soil/pkg/relational"
	"github.com/funtimecoding/soil/pkg/relational/lite"
	"github.com/funtimecoding/soil/pkg/time/clock_tester"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/migrate"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/store"
	"testing"
	"time"
)

func New(t *testing.T) *Tester {
	t.Helper()
	c := clock_tester.NewAt(time.Now().UTC())
	d := lite.NewMemory()
	t.Cleanup(func() { relational.CloseMapper(d) })
	migrate.AutoMigrate(d)

	return &Tester{Store: store.New(d, c.Now), t: t, clock: c}
}
