package example

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/console/status/option"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/system/environment"
	"github.com/funtimecoding/soil/pkg/text/markdown"
	"github.com/funtimecoding/soil/pkg/text/markdown/constant"
	"github.com/funtimecoding/soil/pkg/text/markdown/file"
)

func File() {
	base := environment.Required(constant.WikiPathEnvironment)
	f := option.Color

	for _, n := range system.Files(base) {
		fmt.Printf("File: %s\n", n)
		source := system.ReadBytes(base, n)
		markdown.Print(source, f)
		i := file.New(&source)
		i.Parse()

		if true {
			break
		}
	}
}
