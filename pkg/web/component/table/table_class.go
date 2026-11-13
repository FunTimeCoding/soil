package table

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"maps"
	"slices"
)

func (t *Table) tableClass() string {
	var classes []string

	if t.class != "" {
		classes = append(classes, t.class)
	}

	for _, c := range slices.Sorted(maps.Keys(t.numeric)) {
		classes = append(
			classes,
			fmt.Sprintf("%s%d", constant.NumericColumnClassPrefix, c),
		)
	}

	return join.Space(classes...)
}
