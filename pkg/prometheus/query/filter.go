package query

import (
	"github.com/funtimecoding/soil/pkg/prometheus/query/filter"
	"github.com/funtimecoding/soil/pkg/strings/join/key_value"
)

func Filter(
	name string,
	f *filter.Filter,
) string {
	if f == nil {
		return name
	}

	if f.Count() == 0 {
		return name
	}

	return key_value.Empty(name, f.Build())
}
