package mutation

import "path"

func (r *Report) ByFile() map[string][]*Mutant {
	result := map[string][]*Mutant{}

	for _, f := range r.Files {
		key := path.Join(r.Module, f.Name)
		result[key] = append(result[key], f.Mutants...)
	}

	return result
}
