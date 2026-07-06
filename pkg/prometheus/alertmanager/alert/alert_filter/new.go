package alert_filter

import "github.com/funtimecoding/soil/pkg/prometheus/alertmanager/alert/advanced_option"

func New(o *advanced_option.Alert) *Filter {
	return &Filter{option: o}
}
