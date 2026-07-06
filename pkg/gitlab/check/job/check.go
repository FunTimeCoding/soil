package job

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/console/status/tag"
	"github.com/funtimecoding/soil/pkg/gitlab/check/job/option"
	"github.com/funtimecoding/soil/pkg/gitlab/constant"
	"github.com/funtimecoding/soil/pkg/monitor"
	item "github.com/funtimecoding/soil/pkg/monitor/item/constant"
)

func Check(o *option.Job) {
	elements := collect(o)

	if o.Notation {
		printNotation(elements, o)

		return
	}

	f := constant.CheckFormat

	if o.Copyable {
		f.Tag(tag.Copyable)
	}

	for _, e := range elements {
		fmt.Println(e.Format(f))
	}

	if len(elements) == 0 {
		monitor.NoRelevant(item.GoGitLab.Plural)
	}
}
