package caller

import "example/pkg/target"

func Use(s *target.Shelf) string {
	return s.Name()
}
