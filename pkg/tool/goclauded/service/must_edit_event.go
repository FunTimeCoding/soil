package service

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/event"
)

func (s *Service) MustEditEvent(
	identifier uint,
	message string,
) *event.Event {
	result, e := s.EditEvent(identifier, message)
	errors.PanicOnError(e)

	return result
}
