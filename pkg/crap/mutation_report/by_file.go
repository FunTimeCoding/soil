package mutation_report

import (
	"github.com/funtimecoding/soil/pkg/crap/mutation"
	"path"
)

func (r *Report) ByFile() map[string][]*mutation.Mutant {
	result := map[string][]*mutation.Mutant{}

	for _, f := range r.Files {
		key := path.Join(r.Module, f.Name)
		result[key] = append(result[key], f.Mutants...)
	}

	return result
}
