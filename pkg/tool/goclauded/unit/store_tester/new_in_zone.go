package store_tester

import (
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
	result := &Tester{
		t:     t,
		Store: store.New(lite.NewMemory(), func() time.Time { return now }),
		now:   &now,
	}
	t.Cleanup(result.Store.Close)

	return result
}
