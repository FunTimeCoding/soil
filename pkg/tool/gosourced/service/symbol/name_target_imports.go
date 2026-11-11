package symbol

import (
	sourceConstant "github.com/funtimecoding/soil/pkg/source/constant"
	"github.com/funtimecoding/soil/pkg/source/imports"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/decoration"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/lookup"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/types/relocation"
	"go/ast"
)

func nameTargetImports(
	d *decoration.Set,
	plan *relocation.Plan,
	targetPath string,
	entries []*relocation.Entry,
) {
	file := d.Files[targetPath]
	var existing *ast.File

	if plan.Target != nil {
		existing = lookup.FindSyntaxFile(plan.Set, plan.Target, targetPath)
	}

	v := targetImportView(plan, file, existing, entries)

	for _, decision := range imports.Decide(v) {
		if decision.Kind != sourceConstant.DecisionKeep &&
			decision.Kind != sourceConstant.DecisionRename {
			continue
		}

		alias := decision.Want

		if alias == decision.Entry.Declared {
			alias = ""
		}

		d.AddAlias(file, decision.Entry.Path, alias)
	}
}
