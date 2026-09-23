package service

import "github.com/funtimecoding/soil/pkg/tool/gosproutd/constant"

func (s *Service) Answer(
	identifier uint,
	kind constant.AnswerKind,
	value string,
) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.store.Answer(identifier, kind, constant.AnswerChannelQueue, value)
	s.notifier.Notify()
}
