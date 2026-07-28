package outdated

import (
	"fmt"
	console "github.com/funtimecoding/soil/pkg/console/constant"
	"github.com/funtimecoding/soil/pkg/monitor"
	monitorConstant "github.com/funtimecoding/soil/pkg/monitor/constant"
	"github.com/funtimecoding/soil/pkg/system/macos/brew/constant"
	"github.com/funtimecoding/soil/pkg/system/macos/check/brew/outdated/option"
)

func Check(o *option.Outdated) {
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
		monitor.NoRelevant(monitorConstant.GoBrew.Plural)
	}
}
