package store

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goraidd/store/record"
)

func (s *Store) Raids() []record.Raid {
	var rows []record.Raid
	errors.PanicOnError(
		s.mapper.
			Table("raids").
			Select(
				"raids.id",
				"raids.name",
				"raids.date",
				"count(DISTINCT fights.filename) as fights",
				"count(DISTINCT player_fight_statistics.account) as players",
			).
			Joins("LEFT JOIN fights ON fights.raid_id = raids.id").
			Joins(
				"LEFT JOIN player_fight_statistics ON player_fight_statistics.filename = fights.filename",
			).
			Group("raids.id").
			Order("raids.date DESC").
			Find(&rows).Error,
	)

	return rows
}
