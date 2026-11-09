package transcript

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/context_load"
)

func (s *Service) SessionContextLoads(
	x context.Context,
	sessionIdentifier string,
) ([]context_load.Load, error) {
	loads, e := s.store.SessionContextLoads(sessionIdentifier)

	if e != nil {
		return nil, e
	}

	return s.memory.RedactLoads(x, loads)
}
