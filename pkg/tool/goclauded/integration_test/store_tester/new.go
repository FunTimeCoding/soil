package store_tester

import (
	"github.com/funtimecoding/soil/pkg/relational/lite"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store"
	"testing"
	"time"
)

func New(t *testing.T) *Tester {
	t.Helper()
	now := time.Now()

	return &Tester{
		t: t,
		Store: store.New(
			lite.NewMemory(),
			func() time.Time { return now },
		),
		now: &now,
	}
}
