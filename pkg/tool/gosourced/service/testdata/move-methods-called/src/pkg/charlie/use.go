package charlie

import "example/pkg/alfa"

func Use(s *alfa.Service) string {
	return s.Greet("charlie")
}
