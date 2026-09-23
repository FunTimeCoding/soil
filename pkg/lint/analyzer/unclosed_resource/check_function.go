package unclosed_resource

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/lint/concern"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"go/ast"
	"golang.org/x/tools/go/packages"
)

func checkFunction(
	p *packages.Package,
	results *output.Results,
	body *ast.BlockStmt,
	s *Summaries,
) {
	for _, c := range candidates(p, body) {
		if closes(p, body, c.object) ||
			escapes(p, body, c.object, carriesFrom) ||
			isBorrowed(p, body, c) ||
			s.arrangesFor(calleeOf(p, c.call)) {
			continue
		}

		results.AddConcern(
			concern.NewPosition(
				"unclosed_resource",
				fmt.Sprintf(
					"%s is never closed; use %s",
					c.object.Name(),
					closeSuggestion(c.object),
				),
				p.Fset.Position(c.position),
			),
		)
	}
}
