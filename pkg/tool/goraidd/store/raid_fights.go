package store

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/raid"
)

func (s *Store) RaidFights(raidID int) []raid.Fight {
	var fights []raid.Fight
	errors.PanicOnError(
		s.mapper.
			Where("raid_id = ?", raidID).
			Order("timestamp ASC").
			Find(&fights).Error,
	)

	return fights
}
