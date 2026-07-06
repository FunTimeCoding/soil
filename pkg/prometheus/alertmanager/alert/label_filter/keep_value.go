package label_filter

import "github.com/funtimecoding/soil/pkg/prometheus/alertmanager/alert/label_filter/label"

func (f *Filter) KeepValue(
	k string,
	v string,
) *Filter {
	f.keepValue = append(f.keepValue, label.New(k, v))

	return f
}
