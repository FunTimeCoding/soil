package symbol

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/source/resolve"
	"github.com/funtimecoding/soil/pkg/source/types/resolve_reference"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/lookup"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/commit"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/symbol/identifier"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/sink"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/types/relocation"
	"go/ast"
	"go/token"
	"go/types"
	"golang.org/x/tools/go/packages"
	"maps"
	"path"
	"path/filepath"
)

func (s *Service) moveSymbolsThrough(
	directory string,
	packagePath string,
	symbols []string,
	filePath string,
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

	if filePath != "" && len(symbols) > 0 {
		return commit.FailValidation(r, "pass symbols or file, not both")
	}

	if filePath == "" && len(symbols) == 0 {
		return commit.FailValidation(r, "pass symbols or file")
	}

	all, set, e := s.workspace.ReferenceLoad(
		directory,
		packagePath,
		targetPackagePath,
		func(loaded []*packages.Package, loadedSet *token.FileSet) []types.Object {
			return movedObjects(
				directory,
				loaded,
				loadedSet,
				packagePath,
				symbols,
				filePath,
			)
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

	if filePath != "" {
		full := filePath

		if !filepath.IsAbs(full) {
			full = filepath.Join(directory, filePath)
		}

		file := lookup.FindSyntaxFile(set, p, full)

		if file == nil {
			return commit.FailValidation(
				r,
				fmt.Sprintf("file not found in %s: %s", packagePath, filePath),
			)
		}

		symbols, e = expandFileSymbols(file)

		if e != nil {
			return commit.FailValidation(r, e.Error())
		}
	}

	unique := make(map[string]bool)

	for _, name := range symbols {
		if unique[name] {
			return commit.FailValidation(
				r,
				fmt.Sprintf("duplicate symbol: %s", name),
			)
		}

		unique[name] = true
	}

	entries, message := buildMoveEntries(
		all,
		p,
		packagePath,
		symbols,
		targetFile,
	)

	if message != "" {
		return commit.FailValidation(r, message)
	}

	if message := checkMethodSet(entries); message != "" {
		return commit.FailValidation(r, message)
	}

	references := make(map[*relocation.Entry][]resolve_reference.Reference)

	for _, entry := range entries {
		references[entry] = resolve.FindAllReferences(all, entry.Object)
	}

	flipReferenced(entries, references)
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

	newNames := make(map[string]string)

	for _, entry := range entries {
		if previous, taken := newNames[entry.NewName]; taken {
			return commit.FailValidation(
				r,
				fmt.Sprintf(
					"%s and %s both move as %s",
					previous,
					entry.Symbol,
					entry.NewName,
				),
			)
		}

		newNames[entry.NewName] = entry.Symbol

		if target != nil {
			if f := identifier.CheckScopeCollision(target, entry.NewName); f != nil {
				return commit.FailValidation(r, f.Error())
			}
		}
	}

	flipped, flips, f := planBackReferences(all, p, entries, nil)

	if f != nil {
		return commit.FailValidation(r, f.Error())
	}

	if message := checkEntryGuards(
		all,
		p,
		entries,
		packagePath,
		targetPackagePath,
	); message != "" {
		return commit.FailValidation(r, message)
	}

	moveDirectory, e := lookup.TargetDirectory(p, target, targetPackagePath)

	if e != nil {
		return commit.FailValidation(r, e.Error())
	}

	constraints, message := planConstraints(set, target, entries, moveDirectory)

	if message != "" {
		return commit.FailValidation(r, message)
	}

	targetPackageName := path.Base(targetPackagePath)

	if target != nil {
		targetPackageName = target.Types.Name()
	}

	var external []relocation.QualifiedReference
	renames := make(map[*ast.Ident]string)
	maps.Copy(renames, flips)

	for _, entry := range entries {
		if !entry.Flipped {
			continue
		}

		if capture := embeddedRenames(
			all,
			entry.Object,
			entry.NewName,
			renames,
		); capture != "" {
			return commit.FailValidation(r, capture)
		}
	}

	for _, entry := range entries {
		for _, f := range references[entry] {
			if movedContains(entries, f.Ident.Pos()) {
				if entry.Flipped {
					renames[f.Ident] = entry.NewName
				}

				continue
			}

			external = append(
				external,
				relocation.QualifiedReference{
					Reference: f,
					NewName:   entry.NewName,
				},
			)
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
		return commit.FailValidation(r, blocked)
	}

	reportExports(r, set, flipped)
	resymbolExports(out, set, flipped)
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
