package page

import (
	"bytes"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/yuin/goldmark/v2/parser"
	"github.com/yuin/goldmark/v2/renderer/html"
)

func ToMarkup(markdown string) string {
	s := []byte(markdown)
	var b bytes.Buffer
	errors.PanicOnError(html.New().Render(&b, s, parser.New().Parse(s)))

	return b.String()
}
