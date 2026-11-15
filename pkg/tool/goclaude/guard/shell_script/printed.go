package shell_script

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"mvdan.cc/sh/v3/syntax"
	"strings"
)

func Printed(w *syntax.Word) string {
	var b strings.Builder
	errors.PanicOnError(syntax.NewPrinter().Print(&b, w))

	return b.String()
}
