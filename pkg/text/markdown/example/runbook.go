package example

import (
	"fmt"
	console "github.com/funtimecoding/soil/pkg/console/constant"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/system/environment"
	"github.com/funtimecoding/soil/pkg/text/constant"
	"github.com/funtimecoding/soil/pkg/text/markdown"
	"github.com/funtimecoding/soil/pkg/text/markdown/runbook"
)

func Runbook() {
	base := environment.Required(constant.WikiPathEnvironment)
	f := console.ColorFormat

	for _, n := range system.Files(base) {
		fmt.Printf("File: %s\n", n)
		source := system.ReadBytes(base, n)

		if false {
			markdown.Print(source, f)
		}

		r := runbook.New(&source)
		r.Parse(n)
		fmt.Printf("Runbook: %+v\n", r)
	}
}
