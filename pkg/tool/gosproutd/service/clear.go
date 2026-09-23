package service

import "github.com/funtimecoding/soil/pkg/tool/gosproutd/constant"

func (s *Service) Clear(
	identifier uint,
	line string,
	heard string,
) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if heard != "" && s.store.DecisionState(
		identifier,
	) == constant.StateOpen {
		s.store.Answer(
			identifier,
			constant.AnswerKindOther,
			constant.AnswerChannelConversation,
			heard,
		)
	}

	resolution := constant.ResolutionDefault

	if s.store.DecisionState(identifier) == constant.StateAnswered {
		resolution = constant.ResolutionAnswer
	}

	s.store.Clear(identifier, line, resolution)
	s.notifier.Notify()
}
