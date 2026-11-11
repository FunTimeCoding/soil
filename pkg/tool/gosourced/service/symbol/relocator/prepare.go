package relocator

import (
	"github.com/funtimecoding/soil/pkg/source/resolve"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/lookup"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/symbol/relocation"
	"go/ast"
	"path"
)

func Prepare(
	plan *relocation.Plan,
	skip map[*ast.Ident]bool,
) string {
	flipped, flips, f := planBackReferences(
		plan.All,
		plan.Source,
		plan.Entries,
		skip,
	)

	if f != nil {
		return f.Error()
	}

	if message := checkEntryGuards(
		plan.All,
		plan.Source,
		plan.Entries,
		plan.PackagePath,
		plan.TargetPackagePath,
	); message != "" {
		return message
	}

	moveDirectory, e := lookup.TargetDirectory(
		plan.Source,
		plan.Target,
		plan.TargetPackagePath,
	)

	if e != nil {
		return e.Error()
	}

	constraints, message := planConstraints(
		plan.Set,
		plan.Target,
		plan.Entries,
		moveDirectory,
	)

	if message != "" {
		return message
	}

	plan.Flipped = flipped
	plan.Renames = flips
	plan.MoveDirectory = moveDirectory
	plan.Constraints = constraints
	plan.Resolver = resolve.NewNames(plan.All)
	plan.Qualifications = map[string]*relocation.FileQualification{}
	plan.CreateTarget = plan.Target == nil
	plan.TargetPackageName = path.Base(plan.TargetPackagePath)

	if plan.Target != nil {
		plan.TargetPackageName = plan.Target.Types.Name()
	}

	return ""
}
