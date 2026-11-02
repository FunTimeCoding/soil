package console

import (
	"github.com/funtimecoding/soil/pkg/console/status/option"
	"github.com/funtimecoding/soil/pkg/face"
	stringsConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"strings"
)

func Print(
	v face.Formattable,
	title string,
	indent int,
	f *option.Format,
) {
	Format(
		"%s%s: %s\n",
		strings.Repeat(stringsConstant.DoubleSpace, indent),
		title,
		v.Format(f),
	)
}
