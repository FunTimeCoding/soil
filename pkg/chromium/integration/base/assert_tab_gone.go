package base

import (
	"github.com/funtimecoding/soil/pkg/time/wait_tester"
	"time"
)

func (s *Stack) AssertTabGone(identifier string) {
	s.T.Helper()
	wait_tester.Until(
		s.T,
		func() bool { return !s.TabAlive(identifier) },
		5*time.Second,
	)
}
