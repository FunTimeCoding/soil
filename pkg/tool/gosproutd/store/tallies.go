package store

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/store/decision"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/store/tally"
)

func (s *Store) Tallies() []*tally.Tally {
	var decisions []*decision.Decision
	errors.PanicOnError(
		s.mapper.Preload(
			"Turns",
		).Order(constant.IdentifierColumn).Find(&decisions).Error,
	)
	index := map[string]*tally.Tally{}
	var result []*tally.Tally

	for _, v := range decisions {
		found, okay := index[v.Session]

		if !okay {
			found = tally.New(v.Session)
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
