package store

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/store/decision"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/store/session_statistic"
)

func (s *Store) SessionStatistics() []*session_statistic.Statistic {
	var decisions []*decision.Decision
	errors.PanicOnError(
		s.mapper.Preload(
			"Turns",
		).Order(constant.IdentifierColumn).Find(&decisions).Error,
	)
	index := map[string]*session_statistic.Statistic{}
	var result []*session_statistic.Statistic

	for _, v := range decisions {
		found, okay := index[v.Session]

		if !okay {
			found = session_statistic.New(v.Session)
			index[v.Session] = found
			result = append(result, found)
		}

		found.Emitted++
		found.Bumped += v.Bumped
		found.Turns += len(v.Turns)

		switch v.State {
		case constant.StateAnswered:
			found.Answered++

			if v.AnswerChannel == constant.AnswerChannelConversation {
				found.Bypassed++
			}
		case constant.StateDeclined:
			found.Declined++
		case constant.StateIrrelevant:
			found.Irrelevant++
		case constant.StatePostponed:
			found.Postponed++
		}

		if v.Resolution == constant.ResolutionDefault {
			found.Defaulted++
		}

		if v.State == constant.StateOpen && v.ClearLine == "" {
			found.Abandoned++
		}
	}

	return result
}
