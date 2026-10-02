package goflow

import (
	"errors"
	"github.com/yuin/goldmark/v2/ast"
	"github.com/yuin/goldmark/v2/extension"
	"github.com/yuin/goldmark/v2/parser"
	"slices"
	"strings"
)

func Reflow(
	content string,
	width int,
) (string, error) {
	source := []byte(content)
	body := front(source)
	var patches []*patch
	e := ast.Walk(
		parser.New(parser.WithExtensions(extension.NewTableParser())).
			Parse(source),
		func(
			n ast.Node,
			entering bool,
		) (ast.WalkStatus, error) {
			b, okay := wrappable(n)

			if !entering || !okay || quoted(n) {
				return ast.WalkContinue, nil
			}

			if p := rewrap(source, b, width); p != nil && p.start >= body {
				patches = append(patches, p)
			}

			return ast.WalkContinue, nil
		},
	)

	if e != nil {
		return content, e
	}

	slices.SortFunc(
		patches,
		func(
			a *patch,
			b *patch,
		) int {
			return a.start - b.start
		},
	)
	var b strings.Builder
	at := 0

	for _, p := range patches {
		if p.start < at {
			continue
		}

		b.Write(source[at:p.start])
		b.WriteString(p.content)
		at = p.stop
	}

	b.Write(source[at:])
	rewrapped := b.String()

	if !slices.Equal(words(content), words(rewrapped)) {
		return content, errors.New("word sequence changed")
	}

	return rewrapped, nil
}
