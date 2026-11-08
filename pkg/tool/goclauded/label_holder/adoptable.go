package label_holder

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/types/target"
	"slices"
	"strings"
)

func Adoptable(
	t *target.Target,
	r *Rule,
) bool {
	for _, v := range t.Labels[r.Key] {
		if !slices.ContainsFunc(
			r.Compatible,
			func(c string) bool { return strings.EqualFold(c, v) },
		) {
			return false
		}
	}

	return true
}
