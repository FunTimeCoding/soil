package model_context_tester

import (
	"github.com/funtimecoding/soil/pkg/assert/fixture"
	system "github.com/funtimecoding/soil/pkg/system/constant"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"
)

func (t *Tester) IndexFixtures() {
	t.MustCallTool(
		constant.AddCollection,
		map[string]any{
			"name":        "test",
			constant.Path: fixture.Path(system.SearchPath),
		},
	)
	t.MustCallTool(constant.Index, map[string]any{})
}
