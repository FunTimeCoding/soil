package template

import (
	"bytes"
	"github.com/funtimecoding/soil/pkg/errors"
	"text/template"
)

func Execute(
	t *template.Template,
	a any,
) string {
	var b bytes.Buffer
	errors.PanicOnError(t.Execute(&b, a))

	return b.String()
}
