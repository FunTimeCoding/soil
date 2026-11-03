package gofix

import (
	"github.com/funtimecoding/soil/pkg/lint/output"
	sourceConstant "github.com/funtimecoding/soil/pkg/source/constant"
	"github.com/funtimecoding/soil/pkg/source/imports"
	"go/token"
	"golang.org/x/tools/go/packages"
)

func findImportAliasSuggestions(
	fileSet *token.FileSet,
	all []*packages.Package,
	m *aliasMode,
	r *output.Results,
) ([]Edit, map[token.Pos]bool) {
	var edits []Edit
	handled := make(map[token.Pos]bool)
	seen := make(map[string]bool)

	for _, p := range all {
		for _, file := range p.Syntax {
			if isGenerated(file) {
				continue
			}

			filePath := fileSet.File(file.Pos()).Name()

			if seen[filePath] {
				continue
			}

			seen[filePath] = true
			v, specs := imports.ViewOf(p, file)
			v.Ours = widenOurs(v.Ours, m.ours)
			decisions := imports.Decide(v)

			for _, d := range decisions {
				spec := specs[d.Entry]

				if d.Rule == sourceConstant.RuleConstant {
					handled[spec.Pos()] = true
				}

				if m.suggest {
					reportDecision(d, filePath, r)

					continue
				}

				if d.Rule == sourceConstant.RuleConstant &&
					d.Kind == sourceConstant.DecisionRename {
					reportAliasRename(d, filePath, r)
					edits = append(edits, buildAliasEdits(p, file, spec, d)...)
				}
			}

			if m.suggest {
				continue
			}

			for _, g := range collisionGroups(decisions) {
				if !m.applies(g) {
					continue
				}

				for _, d := range g {
					spec := specs[d.Entry]
					handled[spec.Pos()] = true

					if d.Kind == sourceConstant.DecisionRename {
						reportAliasRename(d, filePath, r)
						edits = append(
							edits,
							buildAliasEdits(p, file, spec, d)...)
					}
				}
			}
		}
	}

	return edits, handled
}
