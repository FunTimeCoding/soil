package store

import "github.com/funtimecoding/soil/pkg/errors"

func (s *Store) RaidPlayerStats(raidID int) []RaidPlayerRow {
	var rows []RaidPlayerRow
	errors.PanicOnError(
		s.mapper.
			Table("player_fight_stats").
			Select(
				"player_fight_stats.account",
				"max(player_fight_stats.name) as name",
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
			Where("fights.raid_id = ?", raidID).
			Group("player_fight_stats.account, profession").
			Order("sum(damage) DESC").
			Find(&rows).Error,
	)

	return rows
}
