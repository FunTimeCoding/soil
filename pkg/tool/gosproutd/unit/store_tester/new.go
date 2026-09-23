package store_tester

import (
	"github.com/funtimecoding/soil/pkg/relational/lite"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/store"
	"testing"
	"time"
)

func New(t *testing.T) *Tester {
	t.Helper()
	now := time.Now().UTC()

	return &Tester{
		Store: store.New(lite.NewMemory(), func() time.Time { return now }),
		t:     t,
		now:   &now,
	}
}
