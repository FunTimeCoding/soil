package report

import (
	soil "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/monitor/collector"
	"github.com/funtimecoding/soil/pkg/monitor/item"
	"time"
)

func (r *Report) AddItem(
	c *collector.Collector,
	identifier string,
	s soil.Severity,
	detail string,
	link string,
	create *time.Time,
) {
	r.Add(item.New(c, identifier, s, detail, link, create))
}
