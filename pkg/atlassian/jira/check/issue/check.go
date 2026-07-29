package issue

import (
	"fmt"
	atlassian "github.com/funtimecoding/soil/pkg/atlassian/constant"
	"github.com/funtimecoding/soil/pkg/atlassian/jira/check/issue/option"
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

	f := atlassian.JiraFormat

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
