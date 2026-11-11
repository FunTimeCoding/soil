package symbol

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/lookup"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/commit"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/symbol/identifier"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/symbol/relocation"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/symbol/relocator"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/sink"
	"go/token"
	"go/types"
	"golang.org/x/tools/go/packages"
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
			return relocator.MethodObjects(
				loaded,
				packagePath,
				typeName,
				methods,
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

	source := relocator.NamedType(p, typeName)

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

	destination := relocator.NamedType(target, targetType)

	if destination == nil {
		return commit.FailValidation(
			r,
			fmt.Sprintf("no named type %s in %s", targetType, targetPackagePath),
		)
	}

	entries, message := relocator.MethodEntries(
		set,
		p,
		source,
		methods,
		targetFile,
	)

	if message != "" {
		return commit.FailValidation(r, message)
	}

	receivers, message := relocator.ReceiverIdentifiers(entries)

	if message != "" {
		return commit.FailValidation(r, message)
	}

	if message := relocator.CheckMethodShape(p, entries, destination); message != "" {
		return commit.FailValidation(r, message)
	}

	if message := relocator.CheckTargetCollision(entries, destination); message != "" {
		return commit.FailValidation(r, message)
	}

	if message := s.checkMovedInterfaces(
		directory,
		all,
		entries,
		relocator.QualifiedType(source),
	); message != "" {
		return commit.FailValidation(r, message)
	}

	holders, message := relocator.RouteCallers(
		all,
		set,
		p,
		source,
		destination,
		entries,
	)

	if message != "" {
		return commit.FailValidation(r, message)
	}

	plan := relocation.NewPlan()
	plan.Set = set
	plan.All = all
	plan.Source = p
	plan.Target = target
	plan.Entries = entries
	plan.Holders = holders
	plan.PackagePath = packagePath
	plan.TargetPackagePath = targetPackagePath

	if message := relocator.Prepare(plan, receivers); message != "" {
		return commit.FailValidation(r, message)
	}

	for i := range receivers {
		plan.Renames[i] = targetType
	}

	for _, entry := range entries {
		out.Resymbol(
			filepath.Dir(set.Position(entry.File.Pos()).Filename),
			identifier.MemberName(typeName, entry.Symbol),
			plan.MoveDirectory,
			identifier.MemberName(targetType, entry.Symbol),
		)
	}

	return relocator.Finish(r, plan, out)
}
