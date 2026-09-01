package alert

import (
	"github.com/funtimecoding/soil/pkg/console"
	consoleConstant "github.com/funtimecoding/soil/pkg/console/constant"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager"
	"github.com/funtimecoding/soil/pkg/prometheus/constant"
	"time"
)

func printRules(
	c *alertmanager.Client,
	firing bool,
) {
	f := consoleConstant.ExtendedColorFormat.Copy()

	for _, r := range c.MustRules().Alert() {
		if r.RawAlert != nil &&
			time.Since(r.RawAlert.LastEvaluation).Round(time.Second) < 1*time.Minute {
			continue
		}

		if firing && r.State != constant.FiringState {
			continue
		}

		console.Line(r.Format(f))
	}
}
