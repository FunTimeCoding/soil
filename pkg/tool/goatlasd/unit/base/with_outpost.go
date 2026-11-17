package base

import (
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/unit/collector_tester"
	outpost "github.com/funtimecoding/soil/pkg/tool/gooutpostd/face"
)

func WithOutpost() Option {
	return func(s *Server) {
		s.Outpost = collector_tester.NewOutpost(
			s.t,
			constant.FixtureVirtualMachineNode,
			nil,
		)
		s.Option.OutpostHosts = []string{constant.FixtureVirtualMachineNode}
		s.Option.Outpost = func(string) outpost.Client { return s.Outpost }
	}
}
