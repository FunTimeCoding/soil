package gocheck

import (
	"github.com/funtimecoding/soil/pkg/lint/output"
	"path"
)

func findingDirectories(sections []*output.Results) map[string]bool {
	result := make(map[string]bool)

	for _, s := range sections {
		for _, c := range s.Entries {
			result[path.Dir(c.Path)] = true
		}
	}

	return result
}
