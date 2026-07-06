package store

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/raid"
	"time"
)

func (s *Store) PlayerAttendance(since time.Time) []AttendanceRow {
	var available int64
	errors.PanicOnError(
		s.mapper.Model(raid.NewRaid()).
			Where("date >= ?", since).
			Count(&available).Error,
	)
	var rows []AttendanceRow
	errors.PanicOnError(
		s.mapper.Model(raid.NewPlayerFightStatistic()).
			Select(
				"account",
				"string_agg(DISTINCT name, ', ' ORDER BY name) as characters",
				"count(DISTINCT fights.raid_id) as fights",
			).
			Joins(
				"JOIN fights ON fights.filename = player_fight_stats.filename",
			).
			Where("fights.raid_id IS NOT NULL").
			Where("fights.timestamp >= ?", since).
			Group("account").
			Order("fights DESC").
			Find(&rows).Error,
	)

	for i := range rows {
		rows[i].Available = int(available)
	}

	return rows
}
