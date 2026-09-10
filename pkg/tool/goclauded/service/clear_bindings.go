package service

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
)

func (s *Service) ClearBindings() {
	for _, e := range s.store.BoundSessions() {
		errors.PanicOnError(
			s.PushQueue(
				e.Identifier,
				e.CallsignValue(),
				constant.QueueReannounce,
				"Re-announce required: MCP binding lost during service restart. Call announce with your session name and topic to restore.",
			),
		)
	}

	s.store.ClearBindings()
}
