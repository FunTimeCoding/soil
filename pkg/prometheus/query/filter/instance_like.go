package filter

import "github.com/funtimecoding/soil/pkg/prometheus/constant"

func (f *Filter) InstanceLike(s string) *Filter {
	return f.Like(constant.InstanceLabel, s)
}
