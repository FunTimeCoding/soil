package item

import (
	"github.com/funtimecoding/soil/pkg/monitor/collector"
	"github.com/funtimecoding/soil/pkg/monitor/constant"
	"time"
)

func New(
	c *collector.Collector,
	identifier string,
	s constant.Severity,
	detail string,
	link string,
	create *time.Time,
) *Item {
	if c == nil {
		panic("nil collector")
	}

	return &Item{
		collector:  c,
		Identifier: identifier,
		Severity:   s,
		Detail:     detail,
		Link:       link,
		Create:     create,
		Status:     constant.Open,
	}
}
