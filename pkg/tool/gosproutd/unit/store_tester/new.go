package store_tester

import (
	"github.com/funtimecoding/soil/pkg/relational"
	"github.com/funtimecoding/soil/pkg/relational/lite"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/migrate"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/store"
	"testing"
	"time"
)

func New(t *testing.T) *Tester {
	t.Helper()
	now := time.Now().UTC()
	d := lite.NewMemory()
	t.Cleanup(func() { relational.CloseMapper(d) })
	migrate.AutoMigrate(d)

	return &Tester{
		Store: store.New(d, func() time.Time { return now }),
		t:     t,
		now:   &now,
	}
}
