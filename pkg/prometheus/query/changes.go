package query

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/prometheus/query/filter"
)

func Changes(
	name string,
	f *filter.Filter,
	minutes int,
) string {
	return fmt.Sprintf("changes(%s[%dm]) > 0", Filter(name, f), minutes)
}
