package base

import (
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/unit/collector_tester"
	"github.com/funtimecoding/soil/pkg/tool/gooutpostd/face"
)

func WithOutpost() Option {
	return func(s *Server) {
		s.Outpost = collector_tester.New(s.t).OutpostClient(
			constant.FixtureVirtualMachineNode,
			nil,
		)
		s.Option.OutpostHosts = []string{constant.FixtureVirtualMachineNode}
		s.Option.Outpost = func(string) face.Client { return s.Outpost }
	}
}
