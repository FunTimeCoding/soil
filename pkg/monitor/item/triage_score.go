package item

import "github.com/funtimecoding/soil/pkg/monitor/constant"

func (i *Item) triageScore() float64 {
	if i.Status == constant.Open && i.Assignee == "" {
		return i.collector.Source.TriageBoost
	}

	return 0.0
}
