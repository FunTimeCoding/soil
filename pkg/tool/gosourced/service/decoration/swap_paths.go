package decoration

import (
	"github.com/dave/dst"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"go/token"
	"strconv"
	"strings"
)

func SwapPaths(
	file *dst.File,
	packagePath string,
	targetPackagePath string,
) {
	prefix := join.Empty(packagePath, "/")
	moved := func(path string) (string, bool) {
		if path != packagePath && !strings.HasPrefix(path, prefix) {
			return path, false
		}

		return join.Empty(
			targetPackagePath,
			strings.TrimPrefix(path, packagePath),
		), true
	}
	dst.Inspect(
		file,
		func(n dst.Node) bool {
			ident, okay := n.(*dst.Ident)

			if !okay {
				return true
			}

			if path, match := moved(ident.Path); match {
				ident.Path = path
			}

			return true
		},
	)

	for _, d := range file.Decls {
		g, okay := d.(*dst.GenDecl)

		if !okay || g.Tok != token.IMPORT {
			continue
		}

		for _, s := range g.Specs {
			spec := s.(*dst.ImportSpec)
			path, e := strconv.Unquote(spec.Path.Value)

			if e != nil {
				continue
			}

			if path, match := moved(path); match {
				spec.Path.Value = strconv.Quote(path)
			}
		}
	}
}
