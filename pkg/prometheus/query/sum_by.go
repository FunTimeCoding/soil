package query

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/prometheus/query/filter"
	"github.com/funtimecoding/soil/pkg/strings/join"
)

func SumBy(
	name string,
	f *filter.Filter,
	by ...string,
) string {
	return fmt.Sprintf(
		"sum(%s) by (%s)",
		Filter(name, f),
		join.CommaSpace(by),
	)
}
