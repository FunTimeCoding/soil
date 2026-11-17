package base

import (
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/constant"
)

func WithFixture(path string) Option {
	return func(s *Server) {
		s.Directory = testutil.PrepareTestPackage(s.t, path)
		s.inventory.Add(constant.FixtureModule, s.Directory)
	}
}
