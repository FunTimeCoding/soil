package base

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/time/wait_tester"
	"strings"
	"time"
)

func (s *Stack) OpenTab(route string) string {
	s.T.Helper()
	identifier, e := s.Client.CreateTab(s.Locator(route))
	assert.FatalOnError(s.T, e)
	title := strings.TrimPrefix(route, "/")
	wait_tester.Until(
		s.T,
		func() bool {
			for _, t := range s.tabs() {
				if t.Identifier == identifier && t.Title == title {
					return true
				}
			}

			return false
		},
		15*time.Second,
	)

	return identifier
}
