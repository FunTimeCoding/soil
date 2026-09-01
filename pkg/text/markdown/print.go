package markdown

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/console/constant"
	"github.com/funtimecoding/soil/pkg/console/status/option"
	"github.com/funtimecoding/soil/pkg/strings"
)

func Print(
	source []byte,
	f *option.Format,
) {
	output := strings.PrefixMultiline(string(source), "> ")

	if f.UseColor {
		output = constant.Cyan("%s", output)
	}

	console.Line(output)
}
