package store_tester

import (
	"github.com/funtimecoding/soil/pkg/relational"
	"github.com/funtimecoding/soil/pkg/relational/lite"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store"
	"testing"
	"time"
)

func NewInZone(
	t *testing.T,
	zone *time.Location,
) *Tester {
	t.Helper()
	now := time.Now().In(zone)
	d := lite.NewMemory()
	t.Cleanup(func() { relational.CloseMapper(d) })

	return &Tester{
		t:     t,
		Store: store.New(d, func() time.Time { return now }),
		now:   &now,
	}
}
