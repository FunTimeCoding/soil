package issue

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/console/age_colorer"
	"github.com/funtimecoding/soil/pkg/console/status/tag"
	"github.com/funtimecoding/soil/pkg/errors/sentry"
	"github.com/funtimecoding/soil/pkg/errors/sentry/check/issue/option"
	"github.com/funtimecoding/soil/pkg/monitor"
	"github.com/funtimecoding/soil/pkg/monitor/item/constant"
)

func Check(o *option.Issue) {
	c := sentry.NewEnvironment()
	elements := c.MustIssuesSimple(o.Verbose)

	if o.Notation {
		printNotation(elements, o)

		return
	}

	f := sentry.Format

	if o.Copyable {
		f.Tag(tag.Copyable)
	}

	colorer := age_colorer.Default(elements...)

	for _, e := range elements {
		colorer.Set(e)
		fmt.Println(e.Format(f))
	}

	if len(elements) == 0 {
		monitor.NoRelevant(constant.GoSentry.Plural)
	}
}
