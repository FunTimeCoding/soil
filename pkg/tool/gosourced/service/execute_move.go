package service

import (
	"fmt"
	"github.com/dave/dst"
	"github.com/funtimecoding/soil/pkg/lint/concern"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/decoration"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/relocation"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/sink"
	"go/ast"
	"sort"
)

func executeMove(
	r *output.Results,
	plan *relocation.Plan,
	out *sink.Sink,
) (*output.Results, error) {
	for _, entry := range plan.Entries {
		if !entry.Flipped {
			continue
		}

		position := plan.Set.Position(entry.Object.Pos())
		r.AddConcern(
			concern.NewLine(
				"exported",
				fmt.Sprintf("%s → %s", entry.Symbol, entry.NewName),
				position.Filename,
				position.Line,
				"",
				true,
			),
		)
	}

	decorations := decoration.NewSet()

	for _, entry := range plan.Entries {
		if _, e := decorations.DecorateFile(
			plan.Set,
			plan.Source,
			entry.File,
		); e != nil {
			return nil, e
		}
	}

	for ident, name := range plan.Renames {
		owner, file := findOwningFile(plan.All, ident.Pos())

		if file == nil {
			continue
		}

		if _, e := decorations.DecorateFile(plan.Set, owner, file); e != nil {
			return nil, e
		}

		if d := decorations.DecoratedIdent(owner, ident); d != nil {
			d.Name = name
		}
	}

	e := applyQualifications(
		r,
		decorations,
		plan.Set,
		plan.Qualifications,
		plan.TargetPackagePath,
	)

	if e != nil {
		return nil, e
	}

	for _, entry := range plan.Entries {
		for _, ident := range entry.BackIdentifiers {
			if d := decorations.DecoratedIdent(plan.Source, ident); d != nil {
				d.Path = plan.PackagePath
			}
		}
	}

	groups := make(map[string][]*relocation.Entry)

	for _, entry := range plan.Entries {
		groups[entry.TargetFile] = append(groups[entry.TargetFile], entry)
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
			plan.Source,
			groups[name],
		)
	}

	removedSpecs := make(map[ast.Spec]bool)
	sourceNames := make(map[string]bool)

	for _, entry := range plan.Entries {
		filename := plan.Set.Position(entry.File.Pos()).Filename
		sourceNames[filename] = true

		if entry.Spec != nil {
			if removedSpecs[entry.Spec] {
				continue
			}

			removedSpecs[entry.Spec] = true
		}

		file := decorations.Files[filename]
		declaration, _ := decorations.DecoratedNode(
			plan.Source,
			entry.Declaration,
		).(dst.Decl)
		var spec dst.Spec

		if entry.Spec != nil {
			spec, _ = decorations.DecoratedNode(
				plan.Source,
				entry.Spec,
			).(dst.Spec)
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

		out.Remove(filename)
		deleted[filename] = true
		r.AddConcern(
			concern.NewFile(
				constant.ConcernRemoved,
				"empty file",
				filename,
				true,
			),
		)
	}

	if plan.CreateTarget {
		out.MakeDirectory(plan.MoveDirectory)
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
						entry.Symbol,
						plan.TargetPackageName,
						entry.NewName,
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
			plan.Resolver,
			decorations.PackagePaths[file],
			decorations.Aliases[file],
			file,
			filename,
			out,
		)

		if e != nil {
			return nil, e
		}
	}

	return r, nil
}
