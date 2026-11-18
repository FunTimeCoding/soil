package service

import (
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/tool/gobatteryd/types/poll"
)

func (s *Service) Read() *poll.Poll {
	result := poll.New()

	for _, t := range s.targets {
		readings, e := readTarget(t)

		if unavailable.Is(e) {
			result.Reachable[t.Server] = false

			continue
		}

		result.Reachable[t.Server] = true

		if e != nil {
			result.Faults = append(result.Faults, e)

			continue
		}

		result.Readings = append(result.Readings, readings...)
	}

	return result
}
