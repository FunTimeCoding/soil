package collector

import (
	"github.com/funtimecoding/soil/pkg/monitor/item/source"
	"time"
)

func New(
	name string,
	prefix string,
	plural string,
	interval time.Duration,
	source *source.Source,
) *Collector {
	return &Collector{
		Name:     name,
		Prefix:   prefix,
		Plural:   plural,
		Interval: interval,
		Source:   source,
	}
}
