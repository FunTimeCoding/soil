package reflow

import (
	"github.com/yuin/goldmark/v2/ast"
	"slices"
	"strings"
)

func Reflow(
	content string,
	width int,
) (string, error) {
	source := []byte(content)
	body := front(source)
	document := parse(source)
	mask := literal(source, document)
	var patches []*patch
	e := ast.Walk(
		document,
		func(
			n ast.Node,
			entering bool,
		) (ast.WalkStatus, error) {
			b, okay := wrappable(n)

			if !entering || !okay || quoted(n) {
				return ast.WalkContinue, nil
			}

			if p := rewrap(source, mask, b, width); p != nil && p.start >= body {
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
	rewrapped := []byte(b.String())

	if f := verify(source, document, rewrapped); f != nil {
		return content, f
	}

	return string(rewrapped), nil
}
