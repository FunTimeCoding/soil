package service

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/context_load"
	"strconv"
)

func (s *Service) redactLoads(loads []context_load.Load) []context_load.Load {
	redacted := s.memory.RedactedMemories()

	if len(redacted) == 0 {
		return loads
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

	return loads
}
