package store

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/session"
)

func (s *Store) AvailableNames() ([]string, error) {
	var taken []string

	if e := s.database.Model(session.Stub()).Where(
		fmt.Sprintf("%s IS NOT NULL", constant.Callsign),
	).Pluck(constant.Callsign, &taken).Error; e != nil {
		return nil, e
	}

	takenSet := map[string]bool{}

	for _, name := range taken {
		takenSet[name] = true
	}

	pool, e := s.poolNames()

	if e != nil {
		return nil, e
	}

	var result []string

	for _, name := range pool {
		if !takenSet[name] {
			result = append(result, name)
		}
	}

	return result, nil
}
