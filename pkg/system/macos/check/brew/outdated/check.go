package outdated

import (
	"fmt"
	consoleConstant "github.com/funtimecoding/soil/pkg/console/constant"
	"github.com/funtimecoding/soil/pkg/monitor"
	monitorConstant "github.com/funtimecoding/soil/pkg/monitor/constant"
	"github.com/funtimecoding/soil/pkg/system/constant"
	"github.com/funtimecoding/soil/pkg/system/macos/check/brew/outdated/option"
)

func Check(o *option.Outdated) {
	elements := collect()

	if o.Notation {
		printNotation(elements, o)

		return
	}

	f := constant.BrewFormat

	if o.Copyable {
		f.Tag(consoleConstant.TagCopyable)
	}

	for _, e := range elements {
		fmt.Println(e.Format(f))
	}

	if len(elements) == 0 {
		monitor.NoRelevant(monitorConstant.GoBrew.Plural)
	}
}
