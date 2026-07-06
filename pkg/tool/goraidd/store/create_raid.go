package store

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/raid"
	"github.com/funtimecoding/soil/pkg/time"
)

func (s *Store) CreateRaid(filenames []string) uint {
	var firstFight raid.Fight
	errors.PanicOnError(
		s.mapper.
			Where("filename IN ?", filenames).
			Order("timestamp ASC").
			First(&firstFight).Error,
	)
	r := raid.NewRaid()
	r.Name = fmt.Sprintf(
		"Raid %s",
		firstFight.Timestamp.Format(time.DateYear),
	)
	r.Date = firstFight.Timestamp
	errors.PanicOnError(s.mapper.Create(r).Error)
	errors.PanicOnError(
		s.mapper.Model(raid.NewFight()).
			Where("filename IN ?", filenames).
			Update("raid_id", r.ID).Error,
	)

	return r.ID
}
