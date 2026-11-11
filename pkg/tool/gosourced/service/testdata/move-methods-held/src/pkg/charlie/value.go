package charlie

import "example/pkg/alfa"

func Value(s alfa.Service) string {
	return s.Greet("value")
}
