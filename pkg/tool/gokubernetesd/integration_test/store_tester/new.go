package store_tester

import (
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/store"
	"path/filepath"
	"testing"
)

func New(t *testing.T) *Tester {
	t.Helper()

	return &Tester{
		t:     t,
		Store: store.New(filepath.Join(t.TempDir(), constant.TestDatabase)),
	}
}
