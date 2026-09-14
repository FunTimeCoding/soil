package page

import (
	"bytes"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/yuin/goldmark/v2/extension"
	"github.com/yuin/goldmark/v2/parser"
	"github.com/yuin/goldmark/v2/renderer/html"
)

func markdownToMarkup(markdown string) string {
	s := []byte(markdown)
	o := parser.New(
		parser.WithExtensions(extension.NewTableParser()),
	).Parse(s)
	var b bytes.Buffer
	errors.PanicOnError(
		html.New(
			html.WithExtensions(extension.NewTableHTMLRenderer()),
		).Render(&b, s, o),
	)

	return b.String()
}
