package file

import (
	"fmt"
	soil "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/monitor/check/file/option"
	"github.com/funtimecoding/soil/pkg/monitor/constant"
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
				constant.GoFile,
				constant.GoFile.IntegerIdentifier(i+1),
				soil.Warning,
				fmt.Sprintf("File old: %s", p),
				"",
				&t,
			)
		} else if o.Verbose {
			r.AddItem(
				constant.GoFile,
				constant.GoFile.IntegerIdentifier(i+1),
				soil.Information,
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
