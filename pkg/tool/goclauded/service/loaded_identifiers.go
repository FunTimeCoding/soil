package service

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"strconv"
)

func (s *Service) loadedIdentifiers(sessionIdentifier string) map[int64]bool {
	loads, e := s.store.ContextLoadsBySession(sessionIdentifier)

	if e != nil {
		return nil
	}

	result := map[int64]bool{}

	for _, entry := range loads {
		if entry.Kind == constant.LoadKindMode {
			continue
		}

		identifier, f := strconv.ParseInt(entry.Reference, 10, 64)

		if f == nil {
			result[identifier] = true
		}
	}

	return result
}
