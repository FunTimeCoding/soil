package store_tester

import (
	"github.com/funtimecoding/soil/pkg/assert/fixture"
	systemConstant "github.com/funtimecoding/soil/pkg/system/constant"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store"
	"testing"
)

func IndexedTestStore(t *testing.T) *store.Store {
	t.Helper()
	s := OpenTestStore(t)
	s.AddCollection(
		"test",
		fixture.Path(systemConstant.SearchPath),
		constant.DefaultGlob,
	)
	s.Index("test")

	return s
}
