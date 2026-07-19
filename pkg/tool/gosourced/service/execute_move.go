package service

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/lint/concern"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"go/ast"
	"os"
	"sort"
)

func executeMove(
	r *output.Results,
	plan *movePlan,
) (*output.Results, error) {
	for _, entry := range plan.entries {
		if !entry.flipped {
			continue
		}

		position := plan.set.Position(entry.object.Pos())
		r.AddConcern(
			concern.NewLine(
				"exported",
				fmt.Sprintf("%s → %s", entry.symbol, entry.newName),
				position.Filename,
				position.Line,
				"",
				true,
			),
		)
	}

	for ident, name := range plan.renames {
		ident.Name = name
	}

	var filenames []string

	for filename := range plan.qualifications {
		filenames = append(filenames, filename)
	}

	sort.Strings(filenames)

	for _, filename := range filenames {
		q := plan.qualifications[filename]
		qualifyReferences(q, plan.packagePath, plan.targetPackagePath)

		for _, qp := range q.positions {
			r.AddConcern(
				concern.NewLine(
					"qualified",
					fmt.Sprintf(
						"%s → %s.%s",
						qp.oldName,
						q.name.local,
						qp.newName,
					),
					qp.position.Filename,
					qp.position.Line,
					"",
					true,
				),
			)
		}
	}

	removedSpecs := make(map[ast.Spec]bool)
	sourceFiles := make(map[string]*ast.File)
	sourceCarried := make(map[string][]*ast.ImportSpec)

	for _, entry := range plan.entries {
		filename := plan.set.Position(entry.file.Pos()).Filename
		sourceFiles[filename] = entry.file
		sourceCarried[filename] = append(
			sourceCarried[filename],
			entry.carried...,
		)

		if entry.spec != nil {
			if removedSpecs[entry.spec] {
				continue
			}

			removedSpecs[entry.spec] = true
		}

		removeDeclaration(entry.file, entry.declaration, entry.spec)
	}

	deleted := make(map[string]bool)
	var sourceNames []string

	for filename := range sourceFiles {
		sourceNames = append(sourceNames, filename)
	}

	sort.Strings(sourceNames)

	for _, filename := range sourceNames {
		file := sourceFiles[filename]

		if hasOnlyImports(file) {
			e := os.Remove(filename)

			if e != nil {
				return nil, e
			}

			deleted[filename] = true
			r.AddConcern(
				concern.NewFile("removed", "empty file", filename, true),
			)

			continue
		}

		removeUnusedImports(
			file,
			sourceCarried[filename],
			collectRemainingImports(file),
		)
	}

	if plan.createTarget {
		e := os.MkdirAll(plan.moveDirectory, 0755)

		if e != nil {
			return nil, e
		}
	}

	for _, entry := range plan.entries {
		if len(entry.backIdentifiers) > 0 {
			rewriteBackReferences(entry, plan.sourceLocalName)
		}
	}

	groups := make(map[string][]*moveEntry)

	for _, entry := range plan.entries {
		groups[entry.targetFile] = append(groups[entry.targetFile], entry)
	}

	var groupNames []string

	for name := range groups {
		groupNames = append(groupNames, name)
	}

	sort.Strings(groupNames)

	for _, name := range groupNames {
		batch := groups[name]
		targetPath, e := writeTargetDeclarations(
			plan.moveDirectory,
			plan.targetPackageName,
			name,
			emitDeclarations(batch),
			uniqueImports(batch),
		)

		if e != nil {
			return nil, e
		}

		for _, entry := range batch {
			r.AddConcern(
				concern.NewFile(
					"moved",
					fmt.Sprintf(
						"%s → %s.%s",
						entry.symbol,
						plan.targetPackageName,
						entry.newName,
					),
					targetPath,
					true,
				),
			)
		}
	}

	for _, filename := range filenames {
		e := writeFile(
			plan.set,
			plan.qualifications[filename].file,
			filename,
		)

		if e != nil {
			return nil, e
		}
	}

	for _, filename := range sourceNames {
		if deleted[filename] {
			continue
		}

		if _, qualified := plan.qualifications[filename]; qualified {
			continue
		}

		e := writeFile(plan.set, sourceFiles[filename], filename)

		if e != nil {
			return nil, e
		}
	}

	return r, nil
}
