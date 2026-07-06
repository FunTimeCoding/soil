package filter

import "github.com/funtimecoding/soil/pkg/prometheus/constant"

func (f *Filter) Node(s string) *Filter {
	return f.Equal(constant.NodeLabel, s)
}
