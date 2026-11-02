package supervisor

import "github.com/funtimecoding/soil/pkg/errors"

func (s *Supervisor) Close() {
	errors.PanicOnError(s.stopAll())
}
