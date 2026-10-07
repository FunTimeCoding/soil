package store

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/raid/model/player_fight_statistic"
	"github.com/funtimecoding/soil/pkg/raid/model/raid"
	"github.com/funtimecoding/soil/pkg/tool/goraidd/store/record"
	"time"
)

func (s *Store) PlayerAttendance(since time.Time) []record.Attendance {
	var available int64
	errors.PanicOnError(
		s.mapper.Model(raid.New()).
			Where("date >= ?", since).
			Count(&available).Error,
	)
	var rows []record.Attendance
	errors.PanicOnError(
		s.mapper.Model(player_fight_statistic.New()).
			Select(
				"account",
				"string_agg(DISTINCT name, ', ' ORDER BY name) as characters",
				"count(DISTINCT fights.raid_id) as fights",
			).
			Joins(
				"JOIN fights ON fights.filename = player_fight_statistics.filename",
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
