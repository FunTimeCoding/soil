package reference

import (
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"github.com/funtimecoding/soil/pkg/lint/pointer"
	"github.com/funtimecoding/soil/pkg/lint/pointer/resolver"
)

func declare(
	bases []string,
	r *resolver.Resolver,
) (*pointer.Declared, []*Finding) {
	result := pointer.NewDeclared()
	var dead []*Finding

	for _, base := range bases {
		if r.BaseExists(base) {
			result.Bases = append(result.Bases, base)

			continue
		}

		dead = append(dead, NewFinding(base, constant.DeadPointerText))
	}

	return result, dead
}
