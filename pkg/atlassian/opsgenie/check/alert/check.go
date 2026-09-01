package alert

import (
	atlassian "github.com/funtimecoding/soil/pkg/atlassian/constant"
	"github.com/funtimecoding/soil/pkg/atlassian/opsgenie/check/alert/option"
	"github.com/funtimecoding/soil/pkg/console"
	consoleConstant "github.com/funtimecoding/soil/pkg/console/constant"
	"github.com/funtimecoding/soil/pkg/monitor"
	monitorConstant "github.com/funtimecoding/soil/pkg/monitor/constant"
)

func Check(o *option.Alert) {
	elements := collect()

	if o.Notation {
		printNotation(elements, o)

		return
	}

	f := atlassian.OpsgenieFormat

	if o.Copyable {
		f.Tag(consoleConstant.TagCopyable)
	}

	for _, e := range elements {
		console.Line(e.Format(f))
	}

	if len(elements) == 0 {
		monitor.NoRelevant(monitorConstant.GoGenie.Plural)
	}
}
