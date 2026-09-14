package service

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/source/resolve"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/decoration"
)

func (s *Service) RenameModule(
	directory string,
	modulePath string,
	newModulePath string,
	force bool,
	dryRun bool,
) (*output.Results, error) {
	r := output.NewResultsWithDirectory(directory)

	if modulePath == newModulePath {
		return failValidation(r, "source and target module are the same")
	}

	if belongsToModule(newModulePath, modulePath) ||
		belongsToModule(modulePath, newModulePath) {
		return failValidation(
			r,
			fmt.Sprintf(
				"%s and %s nest - a module cannot move inside itself",
				modulePath,
				newModulePath,
			),
		)
	}

	all, set, e := loadPackages(directory, "./...")

	if e != nil {
		return nil, e
	}

	symbols := moduleSymbols(all, set, modulePath)

	if len(symbols) == 0 {
		return failValidation(
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
	e = retargetImports(r, decorations, set, all, modulePath, newModulePath)

	if e != nil {
		return nil, e
	}

	names := resolve.NewNames(all)

	for targetPath, target := range targets {
		if target != nil {
			names.Override(targetPath, target.Name())
		}
	}

	e = restoreDecorations(decorations, names, nil, dryRun)

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

	if dryRun {
		r.MarkPlanned()
	}

	return r, nil
}
