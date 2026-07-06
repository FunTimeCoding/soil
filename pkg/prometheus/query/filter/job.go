package filter

import "github.com/funtimecoding/soil/pkg/prometheus/constant"

func (f *Filter) Job(s string) *Filter {
	return f.Equal(constant.JobLabel, s)
}
