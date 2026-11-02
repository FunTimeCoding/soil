package revision

import "path"

func Merge(revisions ...*Revision) *Revision {
	result := &Revision{}

	for _, r := range revisions {
		for _, s := range r.Symbols {
			s.Package = path.Join(r.Module, s.Package)
			result.Symbols = append(result.Symbols, s)
		}

		for _, file := range r.Unparsed {
			result.Unparsed = append(result.Unparsed, path.Join(r.Module, file))
		}
	}

	return result
}
