package report

import (
	"github.com/funtimecoding/soil/pkg/monitor/collector"
	"github.com/funtimecoding/soil/pkg/monitor/constant"
	"github.com/funtimecoding/soil/pkg/monitor/item"
	"time"
)

func (r *Report) AddItem(
	c *collector.Collector,
	identifier string,
	s constant.Severity,
	detail string,
	link string,
	create *time.Time,
) {
	r.Add(item.New(c, identifier, s, detail, link, create))
}
