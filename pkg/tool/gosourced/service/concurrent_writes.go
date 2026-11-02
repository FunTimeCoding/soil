package service

import (
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/constant"
)

func concurrentWrites(r *output.Results) []string {
	var result []string

	for _, c := range r.Entries {
		if c.Key == constant.ConcernConcurrentWrite {
			result = append(result, c.Path)
		}
	}

	return result
}
