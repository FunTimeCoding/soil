package service

import (
	"fmt"
	"github.com/dave/dst"
	"github.com/funtimecoding/soil/pkg/lint/concern"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/decoration"
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

	decorations := decoration.NewSet()

	for _, entry := range plan.entries {
		if _, e := decorations.DecorateFile(plan.set, plan.source, entry.file); e != nil {
			return nil, e
		}
	}

	for ident, name := range plan.renames {
		owner, file := findOwningFile(plan.all, ident.Pos())

		if file == nil {
			continue
		}

		if _, e := decorations.DecorateFile(plan.set, owner, file); e != nil {
			return nil, e
		}

		if d := decorations.DecoratedIdent(owner, ident); d != nil {
			d.Name = name
		}
	}

	var filenames []string

	for filename := range plan.qualifications {
		filenames = append(filenames, filename)
	}

	sort.Strings(filenames)

	for _, filename := range filenames {
		q := plan.qualifications[filename]
		file, e := decorations.DecorateFile(plan.set, q.owner, q.file)

		if e != nil {
			return nil, e
		}

		for ident, newName := range q.idents {
			d := decorations.DecoratedIdent(q.owner, ident)

			if d == nil {
				continue
			}

			d.Name = newName
			d.Path = plan.targetPackagePath
		}

		if q.name != nil && q.name.alias != "" && !q.name.imported {
			decorations.AddAlias(file, plan.targetPackagePath, q.name.alias)
		}

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

	for _, entry := range plan.entries {
		for _, ident := range entry.backIdentifiers {
			if d := decorations.DecoratedIdent(plan.source, ident); d != nil {
				d.Path = plan.packagePath
			}
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
	transplants := make(map[string][]dst.Decl)

	for _, name := range groupNames {
		transplants[name] = transplantEntries(
			decorations,
			plan.source,
			groups[name],
		)
	}

	removedSpecs := make(map[ast.Spec]bool)
	sourceNames := make(map[string]bool)

	for _, entry := range plan.entries {
		filename := plan.set.Position(entry.file.Pos()).Filename
		sourceNames[filename] = true

		if entry.spec != nil {
			if removedSpecs[entry.spec] {
				continue
			}

			removedSpecs[entry.spec] = true
		}

		file := decorations.Files[filename]
		declaration, _ := decorations.DecoratedNode(
			plan.source,
			entry.declaration,
		).(dst.Decl)
		var spec dst.Spec

		if entry.spec != nil {
			spec, _ = decorations.DecoratedNode(plan.source, entry.spec).(dst.Spec)
		}

		removeDecoratedDeclaration(file, declaration, spec)
	}

	deleted := make(map[string]bool)
	var orderedSources []string

	for filename := range sourceNames {
		orderedSources = append(orderedSources, filename)
	}

	sort.Strings(orderedSources)

	for _, filename := range orderedSources {
		if !decoratedHasOnlyImports(decorations.Files[filename]) {
			continue
		}

		if e := os.Remove(filename); e != nil {
			return nil, e
		}

		deleted[filename] = true
		r.AddConcern(concern.NewFile("removed", "empty file", filename, true))
	}

	if plan.createTarget {
		if e := os.MkdirAll(plan.moveDirectory, 0755); e != nil {
			return nil, e
		}
	}

	for _, name := range groupNames {
		targetPath, e := writeMoveTarget(
			decorations,
			plan,
			name,
			transplants[name],
		)

		if e != nil {
			return nil, e
		}

		for _, entry := range groups[name] {
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

	var restoredNames []string

	for filename := range decorations.Files {
		if !deleted[filename] {
			restoredNames = append(restoredNames, filename)
		}
	}

	sort.Strings(restoredNames)

	for _, filename := range restoredNames {
		file := decorations.Files[filename]
		e := restoreDecoratedFile(
			plan.resolver,
			decorations.PackagePaths[file],
			decorations.Aliases[file],
			file,
			filename,
		)

		if e != nil {
			return nil, e
		}
	}

	return r, nil
}
