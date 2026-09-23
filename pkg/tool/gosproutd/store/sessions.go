package store

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/store/decision"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/store/session"
)

func (s *Store) Sessions() []*session.Session {
	var decisions []*decision.Decision
	errors.PanicOnError(
		s.mapper.Order(constant.IdentifierColumn).Find(&decisions).Error,
	)
	index := map[string]*session.Session{}
	var result []*session.Session

	for _, v := range decisions {
		found, okay := index[v.Session]

		if !okay {
			found = session.New(v.Session)
			index[v.Session] = found
			result = append(result, found)
		}

		switch {
		case v.ClearLine != "":
			found.Cleared++
		case v.State == constant.StateOpen:
			found.Open++
		default:
			found.Pending++
		}

		found.Bumped += v.Bumped

		if v.UpdatedAt.After(found.LastAt) {
			found.LastAt = v.UpdatedAt
		}
	}

	return result
}
