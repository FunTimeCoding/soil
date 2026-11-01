package fake_assistant

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
)

func (s *Server) Emit(kind string) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	errors.PanicOnError(
		say(
			s.socket,
			fmt.Sprintf(
				`{"id":%d,"type":"event","event":{"event_type":%q}}`,
				s.identifier,
				kind,
			),
		),
	)
}
