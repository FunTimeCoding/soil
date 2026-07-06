package page

import (
	"bytes"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
)

func markdownToHTML(markdown string) string {
	m := goldmark.New(goldmark.WithExtensions(extension.Table))
	var b bytes.Buffer
	errors.PanicOnError(m.Convert([]byte(markdown), &b))

	return b.String()
}
