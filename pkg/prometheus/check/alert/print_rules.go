package alert

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/console/status/option"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager"
	"github.com/funtimecoding/soil/pkg/prometheus/rule"
	"time"
)

func printRules(
	c *alertmanager.Client,
	firing bool,
) {
	f := option.ExtendedColor.Copy()

	for _, r := range c.MustRules().Alert() {
		if r.RawAlert != nil &&
			time.Since(
				r.RawAlert.LastEvaluation,
			).Round(time.Second) < 1*time.Minute {
			continue
		}

		if firing && r.State != rule.FiringState {
			continue
		}

		fmt.Println(r.Format(f))
	}
}
