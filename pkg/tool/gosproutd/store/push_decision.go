package store

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/store/choice"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/store/decision"
)

func (s *Store) PushDecision(
	session string,
	question string,
	defaultAction string,
	labels []string,
	frames []string,
) *decision.Decision {
	var choices []*choice.Choice

	for i, v := range labels {
		choices = append(choices, choice.New(v, i+1))
	}

	d := decision.New(
		session,
		question,
		defaultAction,
		choices,
		s.resolveFrames(session, frames),
	)
	errors.PanicOnError(s.mapper.Create(d).Error)

	return d
}
