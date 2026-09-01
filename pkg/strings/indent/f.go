package indent

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/strings/constant"
	"strings"
)

func F(
	format string,
	indent int,
	a ...any,
) {
	console.Format(
		"%s%s\n",
		strings.Repeat(constant.DoubleSpace, indent),
		fmt.Sprintf(format, a...),
	)
}
