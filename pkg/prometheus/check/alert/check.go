package alert

import (
	"github.com/funtimecoding/soil/pkg/console"
	consoleConstant "github.com/funtimecoding/soil/pkg/console/constant"
	monitor "github.com/funtimecoding/soil/pkg/monitor/constant"
	"github.com/funtimecoding/soil/pkg/prometheus/check/alert/option"
	"github.com/funtimecoding/soil/pkg/prometheus/constant"
	"github.com/funtimecoding/soil/pkg/tool/common"
)

func Check(o *option.Alert) {
	c := common.Alertmanager()
	alerts, statistic := collect(c, o)

	if o.Notation {
		printNotation(alerts, o)

		return
	}

	if o.Rules {
		printRules(c, o.Firing)

		return
	}

	f := constant.Format.Copy().Tag(consoleConstant.TagHost)

	if o.Copyable {
		f.Tag(consoleConstant.TagCopyable)
	}

	if o.Extended {
		f.Extended()
	}

	m := c.MustRules()

	for _, a := range alerts {
		console.Line(a.Format(f))

		if r := m.Find(a.Name); r != nil {
			console.Format("  Rule: %s\n", r.Format(f))
		}
	}

	if !o.All && statistic.Relevant == 0 {
		console.Format(
			"No relevant %s, %d in total\n",
			monitor.GoAlert.Plural,
			statistic.Total,
		)
	}
}
