package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/context_load"
	"strconv"
)

func (s *Service) redactLoads(
	x context.Context,
	loads []context_load.Load,
) ([]context_load.Load, error) {
	var redacted map[int64]bool

	if e := s.memoryGuard.Do(
		x,
		func(y context.Context) error {
			var f error
			redacted, f = s.memory.RedactedMemories(y)

			return f
		},
	); e != nil {
		return nil, unavailable.Format(constant.MemoryLoadsHidden, e)
	}

	for i := range loads {
		identifier, e := strconv.ParseInt(loads[i].Reference, 10, 64)

		if e != nil || !redacted[identifier] {
			continue
		}

		loads[i].Name = constant.RedactedName
		loads[i].Reference = ""
		loads[i].Query = ""
	}

	return loads, nil
}
