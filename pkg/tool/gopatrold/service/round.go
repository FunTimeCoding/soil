package service

import (
	"context"
	"fmt"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/tool/gopatrold/types/inspection"
)

func (s *Service) Round(x context.Context) ([]*inspection.Inspection, []error) {
	var result []*inspection.Inspection
	var failures []error

	for _, t := range s.targets {
		i := inspection.New(t.Host)
		result = append(result, i)
		arrays, e := peer.Guarded(x, t.Guard, t.Client.Arrays)

		if e != nil {
			failures = append(failures, fmt.Errorf("read arrays: %w", e))

			continue
		}

		i.Reached = true
		i.Arrays = arrays
		drives, f := peer.Guarded(x, t.Guard, t.Client.Drives)

		if f != nil {
			failures = append(failures, fmt.Errorf("read drives: %w", f))

			continue
		}

		i.Drives = drives
	}

	return result, failures
}
