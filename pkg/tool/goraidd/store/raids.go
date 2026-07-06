package store

import "github.com/funtimecoding/soil/pkg/errors"

func (s *Store) Raids() []RaidRow {
	var rows []RaidRow
	errors.PanicOnError(
		s.mapper.
			Table("raids").
			Select(
				"raids.id",
				"raids.name",
				"raids.date",
				"count(DISTINCT fights.filename) as fights",
				"count(DISTINCT player_fight_stats.account) as players",
			).
			Joins("LEFT JOIN fights ON fights.raid_id = raids.id").
			Joins(
				"LEFT JOIN player_fight_stats ON player_fight_stats.filename = fights.filename",
			).
			Group("raids.id").
			Order("raids.date DESC").
			Find(&rows).Error,
	)

	return rows
}
