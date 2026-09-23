package service

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/store/decision"
	"unicode/utf8"
)

func (s *Service) PushDecision(
	session string,
	question string,
	defaultAction string,
	labels []string,
	frames []string,
) (*decision.Decision, error) {
	if length := utf8.RuneCountInString(
		question,
	); length > constant.MaximumTurnLength {
		return nil, fmt.Errorf(
			"question is %d characters, the limit is %d - a question that does not fit is not ready to ask",
			length,
			constant.MaximumTurnLength,
		)
	}

	if defaultAction == "" {
		return nil, fmt.Errorf(
			"a decision needs a default action - what happens if it is never answered",
		)
	}

	s.mutex.Lock()
	defer s.mutex.Unlock()
	d := s.store.PushDecision(session, question, defaultAction, labels, frames)
	s.notifier.Notify()

	return d, nil
}
