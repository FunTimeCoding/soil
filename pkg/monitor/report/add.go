package report

import "github.com/funtimecoding/soil/pkg/monitor/item"

func (r *Report) Add(i *item.Item) {
	i.CalculateScore()
	r.Items = append(r.Items, i)
}
