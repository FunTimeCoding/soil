package issue

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/console/age_colorer"
	console "github.com/funtimecoding/soil/pkg/console/constant"
	"github.com/funtimecoding/soil/pkg/errors/constant"
	"github.com/funtimecoding/soil/pkg/errors/sentry"
	"github.com/funtimecoding/soil/pkg/errors/sentry/check/issue/option"
	"github.com/funtimecoding/soil/pkg/monitor"
	monitorConstant "github.com/funtimecoding/soil/pkg/monitor/constant"
)

func Check(o *option.Issue) {
	c := sentry.NewEnvironment()
	elements := c.MustIssuesSimple(o.Verbose)

	if o.Notation {
		printNotation(elements, o)

		return
	}

	f := constant.Format

	if o.Copyable {
		f.Tag(console.TagCopyable)
	}

	colorer := age_colorer.Default(elements...)

	for _, e := range elements {
		colorer.Set(e)
		fmt.Println(e.Format(f))
	}

	if len(elements) == 0 {
		monitor.NoRelevant(monitorConstant.GoSentry.Plural)
	}
}
