package packages

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/lint/concern"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/source/resolve"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/decoration"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/lookup"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/commit"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/sink"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func (s *Service) movePackageThrough(
	directory string,
	packagePath string,
	targetPackagePath string,
	out *sink.Sink,
) (*output.Results, error) {
	r := output.NewResultsWithDirectory(directory)

	if packagePath == targetPackagePath {
		return commit.FailValidation(
			r,
			"source and target package are the same",
		)
	}

	if path.Base(packagePath) != path.Base(targetPackagePath) {
		return commit.FailValidation(
			r,
			fmt.Sprintf(
				"package name would change from %s to %s - move keeps the name",
				path.Base(packagePath),
				path.Base(targetPackagePath),
			),
		)
	}

	prefix := fmt.Sprintf("%s/", packagePath)

	if strings.HasPrefix(targetPackagePath, prefix) {
		return commit.FailValidation(
			r,
			fmt.Sprintf("cannot move %s into itself", packagePath),
		)
	}

	all, set, e := s.workspace.SubtreeLoad(directory, packagePath)

	if e != nil {
		return nil, e
	}

	p := lookup.FindPackage(all, packagePath)

	if p == nil || len(p.GoFiles) == 0 {
		p = findTestPackage(all, packagePath)
	}

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

	moveDirectory, e := lookup.TargetDirectory(p, nil, targetPackagePath)

	if e != nil {
		return commit.FailValidation(r, e.Error())
	}

	if _, f := os.Stat(moveDirectory); f == nil {
		return commit.FailValidation(
			r,
			fmt.Sprintf("target directory already exists: %s", moveDirectory),
		)
	}

	sourceDirectory := filepath.Dir(p.GoFiles[0])
	decorations := decoration.NewSet()
	e = commit.RetargetImports(
		r,
		decorations,
		set,
		all,
		packagePath,
		targetPackagePath,
	)

	if e != nil {
		return nil, e
	}

	names := resolve.NewNames(all)
	names.Override(targetPackagePath, p.Types.Name())
	e = decoration.Restore(decorations, names, nil, out)

	if e != nil {
		return nil, e
	}

	out.MakeDirectory(filepath.Dir(moveDirectory))
	out.Rename(sourceDirectory, moveDirectory)
	r.AddConcern(
		concern.NewFile(
			constant.ConcernMoved,
			fmt.Sprintf("%s → %s", packagePath, targetPackagePath),
			moveDirectory,
			true,
		),
	)

	return r, nil
}
