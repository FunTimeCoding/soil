package issue

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/atlassian/jira/check/issue/option"
	"github.com/funtimecoding/soil/pkg/atlassian/jira/constant"
	"github.com/funtimecoding/soil/pkg/console/status/tag"
	"github.com/funtimecoding/soil/pkg/monitor"
	item "github.com/funtimecoding/soil/pkg/monitor/item/constant"
)

func Check(o *option.Issue) {
	elements := collect()

	if o.Notation {
		printNotation(elements, o)

		return
	}

	f := constant.Format

	if o.Copyable {
		f.Tag(tag.Copyable)
	}

	for _, e := range elements {
		fmt.Println(e.Format(f))
	}

	if len(elements) == 0 {
		monitor.NoRelevant(item.GoJira.Plural)
	}
}
