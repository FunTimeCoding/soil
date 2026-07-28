package issue

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/atlassian/jira/check/issue/option"
	"github.com/funtimecoding/soil/pkg/atlassian/jira/constant"
	console "github.com/funtimecoding/soil/pkg/console/constant"
	"github.com/funtimecoding/soil/pkg/monitor"
	monitorConstant "github.com/funtimecoding/soil/pkg/monitor/constant"
)

func Check(o *option.Issue) {
	elements := collect()

	if o.Notation {
		printNotation(elements, o)

		return
	}

	f := constant.Format

	if o.Copyable {
		f.Tag(console.TagCopyable)
	}

	for _, e := range elements {
		fmt.Println(e.Format(f))
	}

	if len(elements) == 0 {
		monitor.NoRelevant(monitorConstant.GoJira.Plural)
	}
}
