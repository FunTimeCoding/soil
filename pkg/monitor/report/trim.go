package report

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/monitor/collector"
	"github.com/funtimecoding/soil/pkg/monitor/constant"
	"time"
)

func Trim[T any](
	v []T,
	r *Report,
	all bool,
	o *collector.Collector,
) []T {
	if all {
		return v
	}

	if c := len(v); c > constant.NotationReport {
		v = v[0:constant.NotationReport]
		r.AddItem(
			o,
			o.IntegerIdentifier(0),
			constant.Warning,
			fmt.Sprintf(
				"Too many %s (%d), showing only the newest %d",
				o.Plural,
				c,
				constant.NotationReport,
			),
			"",
			&time.Time{},
		)
	}

	return v
}
