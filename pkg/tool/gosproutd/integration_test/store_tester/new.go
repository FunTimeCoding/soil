package store_tester

import (
	"github.com/funtimecoding/soil/pkg/relational/lite"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/store"
	"testing"
)

func New(t *testing.T) *Tester {
	t.Helper()

	return &Tester{
		Store: store.New(lite.NewMemory()),
	}
}
