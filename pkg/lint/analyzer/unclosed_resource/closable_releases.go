package unclosed_resource

import (
	"github.com/funtimecoding/soil/pkg/lint/fact"
	"go/types"
)

func closableReleases(
	t *types.TypeName,
	releases []fact.Release,
) map[string]bool {
	s, okay := t.Type().Underlying().(*types.Struct)

	if !okay {
		return nil
	}

	result := map[string]bool{}

	for i := range s.NumFields() {
		f := s.Field(i)

		if !carriesObligation(f.Type()) {
			continue
		}

		for _, r := range releases {
			if r.Field == f.Name() {
				result[f.Name()] = true
			}
		}
	}

	return result
}
