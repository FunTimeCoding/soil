package label_filter

import "github.com/funtimecoding/soil/pkg/prometheus/alertmanager/alert/label_filter/label"

func (f *Filter) DropValue(
	k string,
	v string,
) *Filter {
	f.dropValue = append(f.dropValue, label.New(k, v))

	return f
}
