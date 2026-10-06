package service

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/lint/concern"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/source/resolve"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/relocation"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/sink"
	"go/ast"
	"go/token"
	"go/types"
	"golang.org/x/tools/go/packages"
	"path"
	"unicode"
)

func (s *Service) extractTypeThrough(
	directory string,
	packagePath string,
	typeName string,
	targetPackagePath string,
	targetFile string,
	create bool,
	out *sink.Sink,
) (*output.Results, error) {
	r := output.NewResultsWithDirectory(directory)

	if packagePath == targetPackagePath {
		return failValidation(r, "source and target package are the same")
	}

	all, set, e := s.referenceLoad(
		directory,
		packagePath,
		targetPackagePath,
		func(loaded []*packages.Package, _ *token.FileSet) []types.Object {
			p := findPackage(loaded, packagePath)

			if p == nil {
				return nil
			}

			o := p.Types.Scope().Lookup(typeName)

			if o == nil {
				return nil
			}

			return typeMembers(o)
		},
	)

	if e != nil {
		return nil, e
	}

	p := findPackage(all, packagePath)

	if p == nil {
		return failValidation(
			r,
			fmt.Sprintf("package not found: %s", packagePath),
		)
	}

	typeObject := p.Types.Scope().Lookup(typeName)

	if typeObject == nil {
		return failValidation(r, fmt.Sprintf("type not found: %s", typeName))
	}

	if _, okay := typeObject.(*types.TypeName); !okay {
		return failValidation(r, fmt.Sprintf("%s is not a type", typeName))
	}

	entries, message := gatherTypeEntries(set, p, typeObject, targetFile)

	if message != "" {
		return failValidation(r, message)
	}

	target := findPackage(all, targetPackagePath)

	if target == nil && !create {
		return failValidation(
			r,
			fmt.Sprintf(
				"target package not found: %s - pass create to create it",
				targetPackagePath,
			),
		)
	}

	references := make(map[*relocation.Entry][]resolve.Reference)

	for _, entry := range entries {
		references[entry] = resolve.FindAllReferences(all, entry.Object)
	}

	memberNames := typeMemberNames(typeObject)

	for _, entry := range entries {
		if !unicode.IsLower(rune(entry.Symbol[0])) {
			continue
		}

		for _, f := range references[entry] {
			if insideMoved(entries, f.Ident.Pos()) {
				continue
			}

			entry.Flipped = true
			entry.NewName = FlipName(entry.Symbol)

			break
		}

		if !entry.Flipped {
			continue
		}

		if entry.Object != typeObject && memberNames[entry.NewName] {
			return failValidation(
				r,
				fmt.Sprintf(
					"exporting method %s collides with %s",
					entry.Symbol,
					entry.NewName,
				),
			)
		}
	}

	fields, message := fieldFlips(all, entries, typeObject, memberNames)

	if message != "" {
		return failValidation(r, message)
	}

	typeEntry := entries[0]

	if target != nil {
		if f := checkScopeCollision(target, typeEntry.NewName); f != nil {
			return failValidation(r, f.Error())
		}
	}

	if message := checkEntryGuards(
		all,
		p,
		target,
		entries,
		packagePath,
		targetPackagePath,
		false,
	); message != "" {
		return failValidation(r, message)
	}

	moveDirectory, e := targetDirectory(p, target, targetPackagePath)

	if e != nil {
		return failValidation(r, e.Error())
	}

	constraints, message := planConstraints(set, target, entries, moveDirectory)

	if message != "" {
		return failValidation(r, message)
	}

	targetPackageName := path.Base(targetPackagePath)

	if target != nil {
		targetPackageName = target.Types.Name()
	}

	renames := make(map[*ast.Ident]string)
	var external []relocation.QualifiedReference

	for _, entry := range entries {
		isType := entry.Object == typeObject

		for _, f := range references[entry] {
			if insideMoved(entries, f.Ident.Pos()) {
				if entry.Flipped {
					renames[f.Ident] = entry.NewName
				}

				continue
			}

			if isType {
				external = append(
					external,
					relocation.QualifiedReference{
						Reference: f,
						NewName:   entry.NewName,
					},
				)

				continue
			}

			if entry.Flipped {
				renames[f.Ident] = entry.NewName
			}
		}
	}

	for _, field := range fields {
		for _, f := range field.References {
			renames[f.Ident] = field.NewName
		}
	}

	qualifications, blocked := planQualifications(
		set,
		external,
		packagePath,
		targetPackagePath,
		targetPackageName,
	)

	if qualifications == nil {
		return failValidation(
			r,
			fmt.Sprintf("no available import name in %s", blocked),
		)
	}

	for _, field := range fields {
		position := set.Position(field.Object.Pos())
		r.AddConcern(
			concern.NewLine(
				"exported",
				fmt.Sprintf("%s → %s", field.Object.Name(), field.NewName),
				position.Filename,
				position.Line,
				"",
				true,
			),
		)
	}

	plan := relocation.NewPlan()
	plan.Set = set
	plan.All = all
	plan.Source = p
	plan.Target = target
	plan.Resolver = resolve.NewNames(all)
	plan.Entries = entries
	plan.Constraints = constraints
	plan.Qualifications = qualifications
	plan.Renames = renames
	plan.PackagePath = packagePath
	plan.TargetPackagePath = targetPackagePath
	plan.TargetPackageName = targetPackageName
	plan.MoveDirectory = moveDirectory
	plan.CreateTarget = target == nil

	return executeMove(r, plan, out)
}
