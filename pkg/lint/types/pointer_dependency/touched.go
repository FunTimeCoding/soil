package pointer_dependency

import (
	soil "github.com/funtimecoding/soil/pkg/constant"
	"strings"
)

func (s *Set) Touched(changed []string) bool {
	for _, c := range changed {
		if s.paths[c] || s.Tests && strings.HasSuffix(c, soil.TestSuffix) {
			return true
		}

		for p := range s.prefixes {
			if strings.HasPrefix(c, p) {
				return true
			}
		}
	}

	return false
}
