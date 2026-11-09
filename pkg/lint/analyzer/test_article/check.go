package test_article

import (
	"fmt"
	soil "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/lint/concern"
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"go/ast"
	"golang.org/x/tools/go/packages"
	"strings"
)

func Check(
	p *packages.Package,
	results *output.Results,
) {
	for _, file := range p.Syntax {
		if ast.IsGenerated(file) ||
			!strings.HasSuffix(
				p.Fset.Position(file.Pos()).Filename,
				soil.TestSuffix,
			) {
			continue
		}

		for _, d := range file.Decls {
			f, okay := d.(*ast.FuncDecl)

			if !okay || f.Recv != nil ||
				!strings.HasPrefix(f.Name.Name, constant.TestFunctionPrefix) {
				continue
			}

			word := innerArticle(f.Name.Name)

			if word == "" {
				continue
			}

			results.AddConcern(
				concern.NewPosition(
					"test_article",
					fmt.Sprintf(
						"%s carries the article %q - a test name drops a, an and the (see doc/ai/spec/testing.md, Test names)",
						f.Name.Name,
						word,
					),
					p.Fset.Position(f.Name.Pos()),
				),
			)
		}
	}
}
