package service_tester

import (
	"github.com/funtimecoding/soil/pkg/assert/fixture"
	"github.com/funtimecoding/soil/pkg/system/constant"
	goqueryd "github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"
)

func (t *Tester) IndexFixtures() {
	t.Service.AddCollection(
		"test",
		fixture.Path(constant.SearchPath),
		goqueryd.DefaultGlob,
	)
	t.Service.Index("test")
}
