package service_tester

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/relational"
)

func (t *Tester) Close() {
	t.closeOnce.Do(
		func() {
			t.stop.Close()
			t.memoryGuard.Close()
			t.queryGuard.Close()
			t.Set.Close()
			errors.PanicClose(t.searchDatabase)
			relational.CloseMapper(t.Store.Mapper())
		},
	)
}
