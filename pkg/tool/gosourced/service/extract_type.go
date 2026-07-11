package service

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/lint/concern"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/source/resolve"
	"go/ast"
	"go/types"
	"path"
	"unicode"
)

func (s *Service) ExtractType(
	directory string,
	packagePath string,
	typeName string,
	targetPackagePath string,
	targetFile string,
	create bool,
) (*output.Results, error) {
	r := output.NewResultsWithDirectory(directory)

	if packagePath == targetPackagePath {
		return failValidation(r, "source and target package are the same")
	}

	all, set, e := resolve.LoadPackages(directory, "./...")

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
		return failValidation(
			r,
			fmt.Sprintf("type not found: %s", typeName),
		)
	}

	if _, okay := typeObject.(*types.TypeName); !okay {
		return failValidation(
			r,
			fmt.Sprintf("%s is not a type", typeName),
		)
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

	references := make(map[*moveEntry][]resolve.Reference)

	for _, entry := range entries {
		references[entry] = resolve.FindAllReferences(all, entry.object)
	}

	memberNames := typeMemberNames(typeObject)

	for _, entry := range entries {
		if !unicode.IsLower(rune(entry.symbol[0])) {
			continue
		}

		for _, f := range references[entry] {
			if insideMoved(entries, f.Ident.Pos()) {
				continue
			}

			entry.flipped = true
			entry.newName = flipName(entry.symbol)

			break
		}

		if !entry.flipped {
			continue
		}

		if entry.object != typeObject && memberNames[entry.newName] {
			return failValidation(
				r,
				fmt.Sprintf(
					"exporting method %s collides with %s",
					entry.symbol,
					entry.newName,
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
		if f := checkScopeCollision(target, typeEntry.newName); f != nil {
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
	); message != "" {
		return failValidation(r, message)
	}

	moveDirectory, e := targetDirectory(p, target, targetPackagePath)

	if e != nil {
		return failValidation(r, e.Error())
	}

	targetPackageName := path.Base(targetPackagePath)

	if target != nil {
		targetPackageName = target.Types.Name()
	}

	renames := make(map[*ast.Ident]string)
	var external []qualifiedReference

	for _, entry := range entries {
		isType := entry.object == typeObject

		for _, f := range references[entry] {
			if insideMoved(entries, f.Ident.Pos()) {
				if entry.flipped {
					renames[f.Ident] = entry.newName
				}

				continue
			}

			if isType {
				external = append(
					external,
					qualifiedReference{
						reference: f,
						newName:   entry.newName,
					},
				)

				continue
			}

			if entry.flipped {
				renames[f.Ident] = entry.newName
			}
		}
	}

	for _, field := range fields {
		for _, f := range field.references {
			renames[f.Ident] = field.newName
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
		position := set.Position(field.object.Pos())
		r.AddConcern(
			concern.NewLine(
				"exported",
				fmt.Sprintf(
					"%s → %s",
					field.object.Name(),
					field.newName,
				),
				position.Filename,
				position.Line,
				"",
				true,
			),
		)
	}

	return executeMove(
		r,
		&movePlan{
			set:               set,
			entries:           entries,
			qualifications:    qualifications,
			renames:           renames,
			packagePath:       packagePath,
			targetPackagePath: targetPackagePath,
			targetPackageName: targetPackageName,
			moveDirectory:     moveDirectory,
			createTarget:      target == nil,
		},
	)
}
