package store

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/session"
	"math/rand"
)

func (s *Store) NextName() (string, error) {
	var taken []string

	if e := s.database.Model(session.Stub()).Where(
		fmt.Sprintf("%s IS NOT NULL", constant.Callsign),
	).Pluck(constant.Callsign, &taken).Error; e != nil {
		return "", e
	}

	takenSet := map[string]bool{}

	for _, name := range taken {
		takenSet[name] = true
	}

	pool, e := s.poolNames()

	if e != nil {
		return "", e
	}

	var available []string

	for _, name := range pool {
		if !takenSet[name] {
			available = append(available, name)
		}
	}

	if len(available) == 0 {
		return "", nil
	}

	return available[rand.Intn(len(available))], nil
}
