package charlie

import "example/pkg/alfa"

func Hold(s *alfa.Service) func(string) string {
	return s.Greet
}
