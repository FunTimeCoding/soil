package store_tester

import (
	"github.com/funtimecoding/soil/pkg/assert/fixture"
	systemConstant "github.com/funtimecoding/soil/pkg/system/constant"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store"
)

func (o *Tester) IndexedTestStore() *store.Store {
	o.t.Helper()
	s := o.OpenTestStore()
	s.AddCollection(
		"test",
		fixture.Path(systemConstant.SearchPath),
		constant.DefaultGlob,
	)
	s.Index("test")

	return s
}
