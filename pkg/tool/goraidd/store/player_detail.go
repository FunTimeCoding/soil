package store

import "github.com/funtimecoding/soil/pkg/errors"

func (s *Store) PlayerDetail(account string) []PlayerRaidRow {
	var rows []PlayerRaidRow
	errors.PanicOnError(
		s.mapper.
			Table("player_fight_stats").
			Select(
				"raids.id as raid_id",
				"raids.name as raid_name",
				"raids.date as raid_date",
				`(SELECT count(*) FROM fights f2
				  WHERE f2.raid_id = raids.id) as raid_fights`,
				"profession",
				"count(*) as fights",
				"sum(damage) as damage",
				"sum(healing) as healing",
				"sum(condition_cleanses) as condition_cleanses",
				"sum(boon_strips) as boon_strips",
				"sum(barrier) as barrier",
				"sum(downs) as downs",
				"sum(dead_count) as dead_count",
				"sum(active_time_ms) as active_time_ms",
				"avg(dist_to_com) as dist_to_com",
			).
			Joins(
				"JOIN fights ON fights.filename = player_fight_stats.filename",
			).
			Joins("JOIN raids ON raids.id = fights.raid_id").
			Where("player_fight_stats.account = ?", account).
			Group("raids.id, raids.name, raids.date, profession").
			Order("raids.date DESC, sum(damage) DESC").
			Find(&rows).Error,
	)

	return rows
}
