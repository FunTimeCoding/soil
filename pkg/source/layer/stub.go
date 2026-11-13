package layer

import (
	"fmt"
	"go/parser"
	"go/token"
)

func stub(content []byte) []byte {
	name := "ignored"
	f, e := parser.ParseFile(
		token.NewFileSet(),
		"",
		content,
		parser.PackageClauseOnly,
	)

	if e == nil {
		name = f.Name.Name
	}

	return fmt.Appendf(nil, "//go:build ignore\n\npackage %s\n", name)
}
