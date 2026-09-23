package base

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"strings"
	"time"
)

func (s *Stack) OpenTab(route string) string {
	s.T.Helper()
	identifier, e := s.Client.CreateTab(s.Locator(route))
	assert.FatalOnError(s.T, e)
	title := strings.TrimPrefix(route, "/")
	deadline := time.Now().Add(15 * time.Second)

	for time.Now().Before(deadline) {
		for _, t := range s.tabs() {
			if t.Identifier == identifier && t.Title == title {
				return identifier
			}
		}

		time.Sleep(50 * time.Millisecond)
	}

	s.T.Fatalf("tab on %s never reported title %q", route, title)

	return ""
}
