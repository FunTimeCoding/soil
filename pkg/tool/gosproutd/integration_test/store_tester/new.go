package store_tester

import (
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/store"
	"path/filepath"
	"testing"
)

func New(t *testing.T) *Tester {
	t.Helper()

	return &Tester{
		Store: store.New(
			filepath.Join(t.TempDir(), constant.TestDatabase),
		),
	}
}
