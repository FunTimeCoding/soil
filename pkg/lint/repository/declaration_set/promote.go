package declaration_set

import (
	"github.com/funtimecoding/soil/pkg/lint/constant"
	stringsConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
)

func (s *Set) Promote() {
	for owner := range s.Types {
		visited := map[string]bool{}
		pending := []string{owner}

		for len(pending) > 0 {
			current := pending[0]
			pending = pending[1:]

			if visited[current] {
				continue
			}

			visited[current] = true

			for member := range s.members[current] {
				s.Names[join.Empty(owner, stringsConstant.Dot, member)] = true
			}

			if s.foreign[current] {
				s.Names[join.Empty(
					owner,
					stringsConstant.Dot,
					constant.AnyField,
				)] = true
			}

			pending = append(pending, s.embeds[current]...)
		}
	}
}
