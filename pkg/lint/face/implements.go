package face

import "go/types"

func (s *Set) Implements(o types.Object) bool {
	return len(s.Satisfied(o)) > 0
}
