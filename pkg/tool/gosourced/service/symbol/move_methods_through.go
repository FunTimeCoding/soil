package symbol

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/source/resolve"
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
	"path/filepath"
)

func (s *Service) moveMethodsThrough(
	directory string,
	packagePath string,
	typeName string,
	methods []string,
	targetPackagePath string,
	targetType string,
	targetFile string,
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
			return methodObjects(loaded, packagePath, typeName, methods)
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

	source := namedType(p, typeName)

	if source == nil {
		return commit.FailValidation(
			r,
			fmt.Sprintf("no named type %s in %s", typeName, packagePath),
		)
	}

	target := lookup.FindPackage(all, targetPackagePath)

	if target == nil {
		return commit.FailValidation(
			r,
			fmt.Sprintf("target package not found: %s", targetPackagePath),
		)
	}

	destination := namedType(target, targetType)

	if destination == nil {
		return commit.FailValidation(
			r,
			fmt.Sprintf("no named type %s in %s", targetType, targetPackagePath),
		)
	}

	entries, message := methodEntries(set, p, source, methods, targetFile)

	if message != "" {
		return commit.FailValidation(r, message)
	}

	receivers, message := receiverIdentifiers(entries)

	if message != "" {
		return commit.FailValidation(r, message)
	}

	if message := checkMethodShape(p, entries, destination); message != "" {
		return commit.FailValidation(r, message)
	}

	if message := checkTargetCollision(entries, destination); message != "" {
		return commit.FailValidation(r, message)
	}

	if message := s.checkMovedInterfaces(
		directory,
		all,
		entries,
		qualifiedType(source),
	); message != "" {
		return commit.FailValidation(r, message)
	}

	holders, message := routeCallers(all, set, p, source, destination, entries)

	if message != "" {
		return commit.FailValidation(r, message)
	}

	flipped, flips, f := planBackReferences(all, p, entries, receivers)

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

	moveDirectory, g := lookup.TargetDirectory(p, target, targetPackagePath)

	if g != nil {
		return commit.FailValidation(r, g.Error())
	}

	constraints, message := planConstraints(set, target, entries, moveDirectory)

	if message != "" {
		return commit.FailValidation(r, message)
	}

	renames := make(map[*ast.Ident]string)
	maps.Copy(renames, flips)

	for i := range receivers {
		renames[i] = targetType
	}

	for _, o := range flipped {
		if capture := embeddedRenames(
			all,
			o,
			identifier.FlipName(o.Name()),
			renames,
		); capture != "" {
			return commit.FailValidation(r, capture)
		}
	}

	reportExports(r, set, flipped)
	resymbolExports(out, set, flipped)

	for _, entry := range entries {
		out.Resymbol(
			filepath.Dir(set.Position(entry.File.Pos()).Filename),
			identifier.MemberName(typeName, entry.Symbol),
			moveDirectory,
			identifier.MemberName(targetType, entry.Symbol),
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
	plan.Qualifications = map[string]*relocation.FileQualification{}
	plan.Renames = renames
	plan.Holders = holders
	plan.PackagePath = packagePath
	plan.TargetPackagePath = targetPackagePath
	plan.TargetPackageName = target.Types.Name()
	plan.MoveDirectory = moveDirectory

	return executeMove(r, plan, out)
}
