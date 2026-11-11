package symbol

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/source/resolve"
	"github.com/funtimecoding/soil/pkg/source/types/resolve_reference"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/lookup"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/commit"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/symbol/concern_line"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/symbol/identifier"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/symbol/relocation"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/symbol/relocator"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/sink"
	"go/token"
	"go/types"
	"golang.org/x/tools/go/packages"
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
		return commit.FailValidation(
			r,
			"source and target package are the same",
		)
	}

	all, set, e := s.workspace.ReferenceLoad(
		directory,
		packagePath,
		targetPackagePath,
		func(loaded []*packages.Package, _ *token.FileSet) []types.Object {
			p := lookup.FindPackage(loaded, packagePath)

			if p == nil {
				return nil
			}

			o := p.Types.Scope().Lookup(typeName)

			if o == nil {
				return nil
			}

			return relocator.TypeMembers(o)
		},
	)

	if e != nil {
		return nil, e
	}

	p := lookup.FindPackage(all, packagePath)

	if p == nil {
		return commit.FailValidation(
			r,
			fmt.Sprintf("package not found: %s", packagePath),
		)
	}

	typeObject := p.Types.Scope().Lookup(typeName)

	if typeObject == nil {
		return commit.FailValidation(
			r,
			fmt.Sprintf("type not found: %s", typeName),
		)
	}

	if _, okay := typeObject.(*types.TypeName); !okay {
		return commit.FailValidation(
			r,
			fmt.Sprintf("%s is not a type", typeName),
		)
	}

	entries, message := relocator.GatherTypeEntries(
		set,
		p,
		typeObject,
		targetFile,
	)

	if message != "" {
		return commit.FailValidation(r, message)
	}

	target := lookup.FindPackage(all, targetPackagePath)

	if target == nil && !create {
		return commit.FailValidation(
			r,
			fmt.Sprintf(
				"target package not found: %s - pass create to create it",
				targetPackagePath,
			),
		)
	}

	references := make(map[*relocation.Entry][]resolve_reference.Reference)

	for _, entry := range entries {
		references[entry] = resolve.FindAllReferences(all, entry.Object)
	}

	memberNames := relocator.TypeMemberNames(typeObject)
	relocator.FlipReferenced(entries, references)

	for _, entry := range entries {
		if !entry.Flipped {
			continue
		}

		if entry.Object != typeObject && memberNames[entry.NewName] {
			return commit.FailValidation(
				r,
				fmt.Sprintf(
					"exporting method %s collides with %s",
					entry.Symbol,
					entry.NewName,
				),
			)
		}
	}

	fields, message := relocator.FieldFlips(
		all,
		entries,
		typeObject,
		memberNames,
	)

	if message != "" {
		return commit.FailValidation(r, message)
	}

	typeEntry := entries[0]

	if target != nil {
		if f := identifier.CheckScopeCollision(target, typeEntry.NewName); f != nil {
			return commit.FailValidation(r, f.Error())
		}
	}

	plan := relocation.NewPlan()
	plan.Set = set
	plan.All = all
	plan.Source = p
	plan.Target = target
	plan.Entries = entries
	plan.PackagePath = packagePath
	plan.TargetPackagePath = targetPackagePath

	if message := relocator.Prepare(plan, nil); message != "" {
		return commit.FailValidation(r, message)
	}

	if typeEntry.Flipped {
		if capture := relocator.EmbeddedRenames(
			all,
			typeEntry.Object,
			typeEntry.NewName,
			plan.Renames,
		); capture != "" {
			return commit.FailValidation(r, capture)
		}
	}

	var external []relocation.QualifiedReference

	for _, entry := range entries {
		isType := entry.Object == typeObject

		for _, f := range references[entry] {
			if relocator.MovedContains(entries, f.Ident.Pos()) {
				if entry.Flipped {
					plan.Renames[f.Ident] = entry.NewName
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
				plan.Renames[f.Ident] = entry.NewName
			}
		}
	}

	for _, field := range fields {
		for _, f := range field.References {
			plan.Renames[f.Ident] = field.NewName
		}
	}

	qualifications, blocked := relocator.PlanQualifications(
		set,
		external,
		packagePath,
		targetPackagePath,
		plan.TargetPackageName,
	)

	if qualifications == nil {
		return commit.FailValidation(r, blocked)
	}

	plan.Qualifications = qualifications

	for _, field := range fields {
		concern_line.Add(
			r,
			set,
			field.Object.Pos(),
			constant.ConcernExported,
			true,
			"%s → %s",
			field.Object.Name(),
			field.NewName,
		)
	}

	return relocator.Finish(r, plan, out)
}
