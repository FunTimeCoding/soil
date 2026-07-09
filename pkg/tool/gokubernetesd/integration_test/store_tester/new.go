package store_tester

import (
	"github.com/funtimecoding/soil/pkg/relational/lite"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/store"
	"testing"
)

func New(t *testing.T) *Tester {
	t.Helper()

	return &Tester{
		t:     t,
		Store: store.New(lite.NewMemory()),
	}
}
