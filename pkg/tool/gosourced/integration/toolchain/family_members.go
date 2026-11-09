package toolchain

import (
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/constant"
)

func familyMembers(r *output.Results) []string {
	var result []string

	for _, c := range r.Entries {
		if c.Key == constant.ConcernFamily {
			result = append(result, c.Text)
		}
	}

	return result
}
