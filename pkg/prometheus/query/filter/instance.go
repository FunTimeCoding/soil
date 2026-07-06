package filter

import "github.com/funtimecoding/soil/pkg/prometheus/constant"

func (f *Filter) Instance(s string) *Filter {
	return f.Equal(constant.InstanceLabel, s)
}
