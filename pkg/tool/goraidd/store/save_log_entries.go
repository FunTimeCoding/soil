package store

import (
	"github.com/funtimecoding/soil/pkg/gw2/log_manager/log"
	"github.com/funtimecoding/soil/pkg/raid/model/fight"
)

func (s *Store) SaveLogEntries(logs []*log.Log) int {
	result := 0

	for _, l := range logs {
		f := fight.New()
		f.Filename = l.Raw.FileName
		f.Timestamp = l.Time
		f.MapIdentifier = l.Raw.MapIdentifier
		f.AlliedCount = len(l.Raw.Players)

		if r := s.mapper.Where(
			"filename = ?",
			f.Filename,
		).FirstOrCreate(f); r.RowsAffected > 0 {
			result++
		}
	}

	return result
}
