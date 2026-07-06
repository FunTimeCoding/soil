package filter

import "github.com/funtimecoding/soil/pkg/prometheus/constant"

func (f *Filter) Container(s string) *Filter {
	return f.Equal(constant.ContainerLabel, s)
}
