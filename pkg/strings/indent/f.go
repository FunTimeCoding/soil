package indent

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/strings/constant"
	"strings"
)

func F(
	format string,
	indent int,
	a ...any,
) {
	fmt.Printf(
		"%s%s\n",
		strings.Repeat(constant.DoubleSpace, indent),
		fmt.Sprintf(format, a...),
	)
}
