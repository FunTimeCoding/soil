package service_tester

import (
	"github.com/funtimecoding/soil/pkg/assert/fixture"
	systemConstant "github.com/funtimecoding/soil/pkg/system/constant"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"
)

func (t *Tester) IndexFixtures() {
	t.Service.AddCollection(
		"test",
		fixture.Path(systemConstant.SearchPath),
		constant.DefaultGlob,
	)
	t.Service.Index("test")
}
