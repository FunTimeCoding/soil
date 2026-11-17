package base

import (
	"github.com/funtimecoding/soil/pkg/assert/fixture"
	systemConstant "github.com/funtimecoding/soil/pkg/system/constant"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"
)

func (s *Server) IndexFixtures() {
	c := s.Tools()
	c.MustCallTool(
		constant.AddCollection,
		map[string]any{
			"name":        "test",
			constant.Path: fixture.Path(systemConstant.SearchPath),
		},
	)
	c.MustCallTool(constant.Index, map[string]any{})
}
