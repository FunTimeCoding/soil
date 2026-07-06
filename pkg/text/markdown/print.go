package markdown

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/console/status/option"
	"github.com/funtimecoding/soil/pkg/strings"
)

func Print(
	source []byte,
	f *option.Format,
) {
	output := strings.PrefixMultiline(string(source), "> ")

	if f.UseColor {
		output = console.Cyan("%s", output)
	}

	fmt.Println(output)
}
