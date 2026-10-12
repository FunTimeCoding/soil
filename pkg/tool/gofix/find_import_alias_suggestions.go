package gofix

import (
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/source/constant"
	"github.com/funtimecoding/soil/pkg/source/imports"
	"go/token"
	"golang.org/x/tools/go/packages"
)

func findImportAliasSuggestions(
	fileSet *token.FileSet,
	all []*packages.Package,
	realias bool,
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

			for _, d := range imports.Decide(v) {
				spec := specs[d.Entry]

				if d.Rule == constant.RuleConstant {
					handled[spec.Pos()] = true
				}

				reportDecision(d, filePath, realias, r)

				if realias &&
					d.Rule == constant.RuleConstant &&
					d.Kind == constant.DecisionRename {
					edits = append(
						edits,
						buildRealiasEdits(
							p,
							file,
							spec,
							d.Entry.Local,
							d.Want,
							d.Entry.Declared,
						)...,
					)
				}
			}
		}
	}

	return edits, handled
}
