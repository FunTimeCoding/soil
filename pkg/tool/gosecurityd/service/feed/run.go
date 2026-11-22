package feed

import (
	"context"
	"fmt"
	"github.com/funtimecoding/soil/pkg/peer"
	provisionConstant "github.com/funtimecoding/soil/pkg/provision/constant"
	"github.com/funtimecoding/soil/pkg/provision/generated/client"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/constant"
)

func (s *Service) Run(x context.Context) []error {
	run, e := peer.Guarded(
		x,
		s.guard,
		func(y context.Context) (client.HostCommandRun, error) {
			return s.ansible.PostHostCommand(
				y,
				constant.ProbeHostCommand,
				"",
				constant.ProbeSource,
			)
		},
	)

	if e != nil {
		return []error{e}
	}

	var result []error

	for _, h := range run.Hosts {
		if h.Status != provisionConstant.StoreStatusSuccess {
			s.logger.Plain(constant.ProbeSkipped, h.Host, h.Status, h.Stderr)

			continue
		}

		if f := s.store(h.Host, h.Stdout); f != nil {
			result = append(result, fmt.Errorf(constant.ProbeFailed, h.Host, f))
		}
	}

	return result
}
