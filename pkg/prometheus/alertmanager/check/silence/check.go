package silence

import (
	"github.com/funtimecoding/soil/pkg/console"
	consoleConstant "github.com/funtimecoding/soil/pkg/console/constant"
	monitorConstant "github.com/funtimecoding/soil/pkg/monitor/constant"
	"github.com/funtimecoding/soil/pkg/prometheus"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/alert/advanced_option"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/alert_filter"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/check/silence/matcher"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/check/silence/option"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/silence"
	"github.com/funtimecoding/soil/pkg/prometheus/constant"
	"github.com/funtimecoding/soil/pkg/tool/common"
	"time"
)

func Check(o *option.Silence) {
	c := common.Alertmanager()
	defer c.Close()
	silences := collect(c)

	if o.Notation {
		printNotation(silences, o)

		return
	}

	p := prometheus.NewEnvironment()
	defer p.Close()
	r := p.MustRules()

	if o.Set != "" {
		console.Format("Set: %s\n", c.MustSimpleSilence(o.Set, r))
	}

	o2 := advanced_option.New()
	o2.All = true
	a, _ := c.MustAlerts(o2, alert_filter.New(), r)
	console.Format("Alerts: %d\n", len(a))

	if !o.All {
		silences = silence.FilterPermanent(silences)
	}

	var relevant int
	f := constant.AlertmanagerFormat

	if o.Copyable {
		f.Tag(consoleConstant.TagCopyable)
	}

	t := time.Now()

	for _, e := range silences {
		if !o.All && e.State != constant.ActiveState {
			continue
		}

		relevant++
		console.Line(e.Format(f))

		if m := matcher.Matches(e, a, t); len(m) > 0 {
			console.Format("  Matching: %d\n", len(m))
		}
	}

	if !o.All && relevant == 0 {
		console.Format(
			"No relevant %s, %d in total\n",
			monitorConstant.GoSilence.Plural,
			len(silences),
		)
	}
}
