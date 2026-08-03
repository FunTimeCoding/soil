package markdown

import (
	"fmt"
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

	fmt.Println(output)
}
