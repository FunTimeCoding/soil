package alert

import (
	"github.com/funtimecoding/soil/pkg/console"
	consoleConstant "github.com/funtimecoding/soil/pkg/console/constant"
	"github.com/funtimecoding/soil/pkg/prometheus/constant"
	"github.com/funtimecoding/soil/pkg/prometheus/rule/rule_list"
	"time"
)

func printRules(
	l *rule_list.List,
	firing bool,
) {
	f := consoleConstant.ExtendedColorFormat.Copy()

	for _, r := range l.Alert() {
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
