package reference

import (
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"github.com/funtimecoding/soil/pkg/lint/pointer"
)

func declare(
	bases []string,
	r *pointer.Resolver,
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
