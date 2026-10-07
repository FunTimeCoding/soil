package store

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/raid/model/fight"
	"github.com/funtimecoding/soil/pkg/raid/model/raid"
	"github.com/funtimecoding/soil/pkg/time/constant"
)

func (s *Store) CreateRaid(filenames []string) uint {
	var firstFight fight.Fight
	errors.PanicOnError(
		s.mapper.
			Where("filename IN ?", filenames).
			Order("timestamp ASC").
			First(&firstFight).Error,
	)
	r := raid.New()
	r.Name = fmt.Sprintf(
		"Raid %s",
		firstFight.Timestamp.Format(constant.DateYear),
	)
	r.Date = firstFight.Timestamp
	errors.PanicOnError(s.mapper.Create(r).Error)
	errors.PanicOnError(
		s.mapper.Model(fight.New()).
			Where("filename IN ?", filenames).
			Update("raid_id", r.Identifier).Error,
	)

	return r.Identifier
}
