package preposition_name

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/lint/concern"
	"github.com/funtimecoding/soil/pkg/lint/face"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"go/ast"
	"golang.org/x/tools/go/packages"
)

func checkName(
	p *packages.Package,
	results *output.Results,
	faces *face.Set,
	ident *ast.Ident,
) {
	o := p.TypesInfo.Defs[ident]

	if o == nil {
		return
	}

	word := spokenWord(o)

	if word == "" || exempt(o, faces, word) {
		return
	}

	results.AddConcern(
		concern.NewPosition(
			"preposition_name",
			fmt.Sprintf(
				"%s carries %q - a preposition stands where the owner went missing (see doc/ai/spec/naming.md, Prepositions)",
				ident.Name,
				word,
			),
			p.Fset.Position(ident.Pos()),
		),
	)
}
