package store

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/raid/model/fight"
)

func (s *Store) RaidFights(raidIdentifier int) []fight.Fight {
	var fights []fight.Fight
	errors.PanicOnError(
		s.mapper.
			Where("raid_id = ?", raidIdentifier).
			Order("timestamp ASC").
			Find(&fights).Error,
	)

	return fights
}
