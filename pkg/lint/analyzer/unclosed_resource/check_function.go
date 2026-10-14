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
		if closes(p, body, c.Object, s) ||
			escapes(p, body, c.Object, carriesFrom) ||
			s.isUnowned(p, body, c) {
			continue
		}

		results.AddConcern(
			concern.NewPosition(
				"unclosed_resource",
				fmt.Sprintf(
					"%s is never closed; use %s",
					c.Object.Name(),
					closeSuggestion(c.Object),
				),
				p.Fset.Position(c.Position),
			),
		)
	}

	owned := ownedLocals(p, body, s)

	for _, t := range transfers(p, body) {
		if ownsValue(p, body, s, owned, t.Value) &&
			!fieldReleasedIn(p, body, t, s) {
			checkTransfer(p, results, s, t)
		}
	}
}
