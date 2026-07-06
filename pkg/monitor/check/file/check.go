package file

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/monitor/check/file/option"
	"github.com/funtimecoding/soil/pkg/monitor/constant"
	item "github.com/funtimecoding/soil/pkg/monitor/item/constant"
	"github.com/funtimecoding/soil/pkg/monitor/report"
	"github.com/funtimecoding/soil/pkg/system"
	"time"
)

func Check(o *option.File) {
	r := report.New()

	if o.Verbose {
		r.AddNote("verbose output")
	}

	for i, p := range collect(o.Paths) {
		t := system.Stat(p).ModTime()

		if time.Since(t) > 5*time.Minute {
			r.AddItem(
				item.GoFile,
				item.GoFile.IntegerIdentifier(i+1),
				constant.Warning,
				fmt.Sprintf("File old: %s", p),
				"",
				&t,
			)
		} else if o.Verbose {
			r.AddItem(
				item.GoFile,
				item.GoFile.IntegerIdentifier(i+1),
				constant.Information,
				fmt.Sprintf("File good: %s", p),
				"",
				&t,
			)
		} else {
			r.AddNote("good: %s", p)
		}
	}

	r.Print()
}
