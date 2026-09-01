package example

import (
	"github.com/funtimecoding/soil/pkg/console"
	consoleConstant "github.com/funtimecoding/soil/pkg/console/constant"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/system/environment"
	"github.com/funtimecoding/soil/pkg/text/constant"
	"github.com/funtimecoding/soil/pkg/text/markdown"
	"github.com/funtimecoding/soil/pkg/text/markdown/runbook"
)

func Runbook() {
	base := environment.Required(constant.WikiPathEnvironment)
	f := consoleConstant.ColorFormat

	for _, n := range system.Files(base) {
		console.Format("File: %s\n", n)
		source := system.ReadBytes(base, n)

		if false {
			markdown.Print(source, f)
		}

		r := runbook.New(&source)
		r.Parse(n)
		console.Format("Runbook: %+v\n", r)
	}
}
