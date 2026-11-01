package store_tester

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/relational/lite/connection"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/store"
	"testing"
)

func New(t *testing.T) *Tester {
	t.Helper()
	d := connection.NewMemory()
	t.Cleanup(func() { errors.PanicClose(d) })

	return &Tester{t: t, Store: store.New(d)}
}
