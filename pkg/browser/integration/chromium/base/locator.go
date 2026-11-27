package base

import "github.com/funtimecoding/soil/pkg/strings/join"

func (s *Stack) Locator(route string) string {
	return join.Empty(s.Site.URL, route)
}
