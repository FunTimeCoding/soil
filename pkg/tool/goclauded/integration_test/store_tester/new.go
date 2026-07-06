package store_tester

import (
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store"
	"path/filepath"
	"testing"
	"time"
)

func New(t *testing.T) *Tester {
	t.Helper()
	now := time.Now()

	return &Tester{
		t: t,
		Store: store.New(
			filepath.Join(t.TempDir(), constant.TestDatabase),
			func() time.Time { return now },
		),
		now: &now,
	}
}
