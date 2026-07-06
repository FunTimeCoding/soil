package console

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/console/status/option"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/strings/separator"
	"strings"
)

func Print(
	v face.Formattable,
	title string,
	indent int,
	f *option.Format,
) {
	fmt.Printf(
		"%s%s: %s\n",
		strings.Repeat(separator.DoubleSpace, indent),
		title,
		v.Format(f),
	)
}
