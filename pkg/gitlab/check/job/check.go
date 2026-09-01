package job

import (
	"github.com/funtimecoding/soil/pkg/console"
	consoleConstant "github.com/funtimecoding/soil/pkg/console/constant"
	"github.com/funtimecoding/soil/pkg/gitlab/check/job/option"
	"github.com/funtimecoding/soil/pkg/gitlab/constant"
	"github.com/funtimecoding/soil/pkg/monitor"
	monitorConstant "github.com/funtimecoding/soil/pkg/monitor/constant"
)

func Check(o *option.Job) {
	elements := collect(o)

	if o.Notation {
		printNotation(elements, o)

		return
	}

	f := constant.CheckFormat

	if o.Copyable {
		f.Tag(consoleConstant.TagCopyable)
	}

	for _, e := range elements {
		console.Line(e.Format(f))
	}

	if len(elements) == 0 {
		monitor.NoRelevant(monitorConstant.GoGitLab.Plural)
	}
}
