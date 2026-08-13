package silence

import (
	"github.com/funtimecoding/soil/pkg/constant"
	monitor "github.com/funtimecoding/soil/pkg/monitor/constant"
	"github.com/funtimecoding/soil/pkg/monitor/report"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/check/silence/option"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/silence"
)

func printNotation(
	v []*silence.Silence,
	o *option.Silence,
) {
	r := report.New()

	for _, e := range report.Trim(v, r, o.All, monitor.GoSilence) {
		r.AddItem(
			monitor.GoSilence,
			e.MonitorIdentifier,
			constant.Warning,
			e.Rule,
			e.Link,
			e.Start,
		)
	}

	r.Print()
}
