package store

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goraidd/store/record"
)

func (s *Store) RaidPlayerStats(raidIdentifier int) []record.RaidPlayer {
	var rows []record.RaidPlayer
	errors.PanicOnError(
		s.mapper.
			Table("player_fight_statistics").
			Select(
				"player_fight_statistics.account",
				"max(player_fight_statistics.name) as name",
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
				"JOIN fights ON fights.filename = player_fight_statistics.filename",
			).
			Where("fights.raid_id = ?", raidIdentifier).
			Group("player_fight_statistics.account, profession").
			Order("sum(damage) DESC").
			Find(&rows).Error,
	)

	return rows
}
