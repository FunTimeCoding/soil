package label_filter

import "github.com/funtimecoding/soil/pkg/prometheus/alertmanager/alert/label_filter/label"

type Filter struct {
	keep          []string
	drop          []string
	keepValue     []*label.Label
	dropValue     []*label.Label
	KeepByDefault bool
}
