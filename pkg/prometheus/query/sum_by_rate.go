package query

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/prometheus/query/filter"
	"github.com/funtimecoding/soil/pkg/strings/join"
)

func SumByRate(
	name string,
	f *filter.Filter,
	minutes int,
	by ...string,
) string {
	return fmt.Sprintf(
		"sum(%s) by (%s)",
		Rate(name, f, minutes),
		join.CommaSpace(by),
	)
}
