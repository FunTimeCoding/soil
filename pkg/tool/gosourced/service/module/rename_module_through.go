package module

import (
	"fmt"
	goModConstant "github.com/funtimecoding/soil/pkg/go_mod/constant"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/source/resolve"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/decoration"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/load"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/commit"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/sink"
)

func renameModuleThrough(
	directory string,
	modulePath string,
	newModulePath string,
	force bool,
	dryRun bool,
	out *sink.Sink,
) (*output.Results, error) {
	r := output.NewResultsWithDirectory(directory)

	if modulePath == newModulePath {
		return commit.FailValidation(r, "source and target module are the same")
	}

	if belongsToModule(newModulePath, modulePath) ||
		belongsToModule(modulePath, newModulePath) {
		return commit.FailValidation(
			r,
			fmt.Sprintf(
				"%s and %s nest - a module cannot move inside itself",
				modulePath,
				newModulePath,
			),
		)
	}

	all, set, e := load.Packages(directory, goModConstant.AllPackages)

	if e != nil {
		return nil, e
	}

	symbols := moduleSymbols(all, set, modulePath)

	if len(symbols) == 0 {
		return commit.FailValidation(
			r,
			fmt.Sprintf("no references to %s in this module", modulePath),
		)
	}

	targets := moduleTargetPackages(
		directory,
		symbols,
		modulePath,
		newModulePath,
	)
	breakages := moduleBreakages(targets, symbols, modulePath, newModulePath)

	if len(breakages) > 0 && !force && !dryRun {
		for _, c := range breakages {
			r.AddConcern(c)
		}

		return r, nil
	}

	decorations := decoration.NewSet()
	e = commit.RetargetImports(
		r,
		decorations,
		set,
		all,
		modulePath,
		newModulePath,
	)

	if e != nil {
		return nil, e
	}

	names := resolve.NewNames(all)

	for targetPath, target := range targets {
		if target != nil {
			names.Override(targetPath, target.Name())
		}
	}

	e = decoration.Restore(decorations, names, nil, out)

	if e != nil {
		return nil, e
	}

	for _, c := range breakages {
		c.Fixed = true
		r.AddConcern(c)
	}

	r.AddConcern(
		moduleRenamed(modulePath, newModulePath, len(symbols), len(breakages)),
	)

	return r, nil
}
