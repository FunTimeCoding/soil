package example

import (
	"github.com/funtimecoding/soil/pkg/console"
	consoleConstant "github.com/funtimecoding/soil/pkg/console/constant"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/system/environment"
	"github.com/funtimecoding/soil/pkg/text/constant"
	"github.com/funtimecoding/soil/pkg/text/markdown"
	"github.com/funtimecoding/soil/pkg/text/markdown/file"
)

func File() {
	base := environment.Required(constant.WikiPathEnvironment)
	f := consoleConstant.ColorFormat

	for _, n := range system.Files(base) {
		console.Format("File: %s\n", n)
		source := system.ReadBytes(base, n)
		markdown.Print(source, f)
		i := file.New(&source)
		i.Parse()

		if true {
			break
		}
	}
}
