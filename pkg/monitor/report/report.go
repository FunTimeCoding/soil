package report

import "github.com/funtimecoding/soil/pkg/monitor/item"

type Report struct {
	Count int
	Items []*item.Item
	Note  []string
}
