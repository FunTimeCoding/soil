package item

import soil "github.com/funtimecoding/soil/pkg/constant"

func (i *Item) triageScore() float64 {
	if i.Status == soil.Open && i.Assignee == "" {
		return i.collector.Source.TriageBoost
	}

	return 0.0
}
