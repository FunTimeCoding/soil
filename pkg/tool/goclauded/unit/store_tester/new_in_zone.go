package store_tester

import (
	"github.com/funtimecoding/soil/pkg/relational"
	"github.com/funtimecoding/soil/pkg/relational/lite"
	"github.com/funtimecoding/soil/pkg/time/clock_tester"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/migrate"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store"
	"testing"
	"time"
)

func NewInZone(
	t *testing.T,
	zone *time.Location,
) *Tester {
	t.Helper()
	c := clock_tester.NewAt(time.Now().In(zone))
	d := lite.NewMemory()
	t.Cleanup(func() { relational.CloseMapper(d) })
	migrate.AutoMigrate(d)

	return &Tester{t: t, mapper: d, Store: store.New(d, c.Now), clock: c}
}
