package service

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"go/ast"
	"go/token"
	"sort"
	"strings"
)

func statementShape(
	content []byte,
	set *token.FileSet,
	node ast.Node,
	anchor ast.Node,
) (string, string) {
	start := set.Position(node.Pos()).Offset
	end := set.Position(node.End()).Offset
	lineEnd := end

	if index := strings.IndexByte(string(content[start:end]), '\n'); index >= 0 {
		lineEnd = start + index
	}

	type replacement struct {
		start int
		end   int
		text  string
	}
	var replacements []replacement
	ast.Inspect(
		node,
		func(n ast.Node) bool {
			if n == nil {
				return false
			}

			if n.Pos() >= anchor.Pos() && n.End() <= anchor.End() {
				return false
			}

			from := set.Position(n.Pos()).Offset
			to := set.Position(n.End()).Offset

			if to > lineEnd {
				return true
			}

			switch leaf := n.(type) {
			case *ast.Ident:
				replacements = append(
					replacements,
					replacement{start: from, end: to, text: "IDENT"},
				)
			case *ast.BasicLit:
				replacements = append(
					replacements,
					replacement{start: from, end: to, text: leaf.Kind.String()},
				)
			}

			return true
		},
	)
	sort.Slice(
		replacements,
		func(i, j int) bool {
			return replacements[i].start > replacements[j].start
		},
	)
	line := string(content[start:lineEnd])
	exemplar := strings.TrimSpace(line)

	for _, r := range replacements {
		line = join.Empty(line[:r.start-start], r.text, line[r.end-start:])
	}

	return strings.TrimSpace(line), exemplar
}
