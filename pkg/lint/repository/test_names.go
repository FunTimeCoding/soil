package repository

import (
	soil "github.com/funtimecoding/soil/pkg/constant"
	"strings"
)

func (r *Repository) testNames() map[string]map[string]bool {
	result := map[string]map[string]bool{}

	for _, path := range r.Files.Files() {
		if strings.HasSuffix(path, soil.TestSuffix) && !r.Ignored(path) {
			addTestNames(result, r.Absolute(path))
		}
	}

	for _, sibling := range r.Siblings {
		for _, path := range siblingTests(
			r.Absolute(sibling),
			r.siblingIgnored[sibling],
		) {
			addTestNames(result, path)
		}
	}

	return result
}
