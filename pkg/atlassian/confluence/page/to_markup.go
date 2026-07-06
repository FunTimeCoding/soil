package page

import (
	"bytes"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/yuin/goldmark"
)

func ToMarkup(markdown string) string {
	var b bytes.Buffer
	errors.PanicOnError(goldmark.Convert([]byte(markdown), &b))

	return b.String()
}
