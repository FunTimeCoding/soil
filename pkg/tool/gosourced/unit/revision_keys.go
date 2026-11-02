package unit

import "github.com/funtimecoding/soil/pkg/tool/gosourced/revision"

func revisionKeys(r *revision.Revision) []string {
	var result []string

	for _, s := range r.Symbols {
		result = append(result, s.Key())
	}

	return result
}
