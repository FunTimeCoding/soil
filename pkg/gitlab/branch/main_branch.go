package branch

import (
	"github.com/funtimecoding/soil/pkg/git/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"log"
	"slices"
)

func MainBranch(v []*Branch) *Branch {
	var names []string

	for _, b := range v {
		if !slices.Contains(names, b.Name) {
			names = append(names, b.Name)
		}

		if slices.Contains(constant.MainBranches, b.Name) {
			return b
		}
	}

	log.Panicf("main branch not found: %s", join.Comma(names))

	return &Branch{}
}
