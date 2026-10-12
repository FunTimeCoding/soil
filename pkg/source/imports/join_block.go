package imports

import (
	"bytes"
	"go/ast"
	"go/parser"
	"go/token"
	"sort"
	"strings"
)

func JoinBlock(source []byte) []byte {
	fileSet := token.NewFileSet()
	file, e := parser.ParseFile(
		fileSet,
		"",
		source,
		parser.ImportsOnly|parser.ParseComments,
	)

	if e != nil {
		return source
	}

	for _, d := range file.Decls {
		g, okay := d.(*ast.GenDecl)

		if !okay || g.Tok != token.IMPORT || !g.Lparen.IsValid() {
			continue
		}

		start := fileSet.Position(g.Lparen).Offset + 1
		end := fileSet.Position(g.Rparen).Offset
		block := string(source[start:end])

		if !strings.Contains(block, "\n\n") || strings.Contains(block, "//") ||
			strings.Contains(block, "/*") {
			continue
		}

		var lines []string

		for _, line := range strings.Split(block, "\n") {
			if trimmed := strings.TrimSpace(line); trimmed != "" {
				lines = append(lines, trimmed)
			}
		}

		sort.Slice(
			lines,
			func(i, j int) bool { return specPath(lines[i]) < specPath(lines[j]) },
		)
		var b bytes.Buffer
		b.Write(source[:start])
		b.WriteString("\n")

		for _, line := range lines {
			b.WriteString("\t")
			b.WriteString(line)
			b.WriteString("\n")
		}

		b.Write(source[end:])

		return b.Bytes()
	}

	return source
}
