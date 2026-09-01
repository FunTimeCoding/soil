package console

import (
	"github.com/funtimecoding/soil/pkg/console/status/option"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/strings/constant"
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
		strings.Repeat(constant.DoubleSpace, indent),
		title,
		v.Format(f),
	)
}
