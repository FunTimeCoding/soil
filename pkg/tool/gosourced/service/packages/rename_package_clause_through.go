package packages

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/source/resolve"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/decoration"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/lookup"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/commit"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/sink"
	"go/ast"
	"go/token"
	"path/filepath"
)

func (s *Service) renamePackageClauseThrough(
	directory string,
	packagePath string,
	newName string,
	out *sink.Sink,
) (*output.Results, error) {
	r := output.NewResultsWithDirectory(directory)

	if !token.IsIdentifier(newName) {
		return commit.FailValidation(
			r,
			fmt.Sprintf("not a valid package name: %s", newName),
		)
	}

	all, set, e := s.workspace.SubtreeLoad(directory, packagePath)

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

	if len(p.GoFiles) == 0 {
		return commit.FailValidation(
			r,
			fmt.Sprintf("package has no Go files: %s", packagePath),
		)
	}

	oldName := p.Types.Name()

	if oldName == newName {
		return commit.FailValidation(
			r,
			fmt.Sprintf("package is already named %s", newName),
		)
	}

	qualifiers, taken := collectPackageQualifiers(
		all,
		set,
		packagePath,
		oldName,
		newName,
	)

	if taken != "" {
		return commit.FailValidation(r, taken)
	}

	modified := make(map[string]*ast.File)
	renamePackageClauses(
		all,
		set,
		filepath.Dir(p.GoFiles[0]),
		oldName,
		modified,
	)
	decorations := decoration.NewSet()

	if e := decorateModified(decorations, set, all, modified); e != nil {
		return nil, e
	}

	reportClauses(
		r,
		set,
		modified,
		filepath.Dir(p.GoFiles[0]),
		oldName,
		newName,
	)
	renameDecoratedClauses(decorations, modified, oldName, newName)
	e = decorateQualifiers(r, decorations, set, qualifiers, oldName, newName)

	if e != nil {
		return nil, e
	}

	names := resolve.NewNames(all)
	names.Override(packagePath, newName)

	if e := decoration.Restore(decorations, names, nil, out); e != nil {
		return nil, e
	}

	return r, nil
}
