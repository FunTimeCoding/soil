package item

import (
	soil "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/monitor/collector"
	"time"
)

type Item struct {
	collector  *collector.Collector
	Identifier string
	Label      string
	Severity   soil.Severity
	Detail     string
	Status     soil.Status
	Assignee   string
	Link       string
	Create     *time.Time
	Update     *time.Time
	Score      float64
}
