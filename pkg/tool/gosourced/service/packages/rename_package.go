package packages

import (
	"context"
	go_modConstant "github.com/funtimecoding/soil/pkg/go_mod/constant"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/sink"
)

func (s *Service) RenamePackage(
	x context.Context,
	directory string,
	packagePath string,
	newName string,
	dryRun bool,
) (*output.Results, error) {
	return s.commit.Transact(
		x,
		directory,
		[]string{join.Empty(packagePath, go_modConstant.RecursivePattern)},
		dryRun,
		func(d string, out *sink.Sink) (*output.Results, error) {
			if d != directory {
				return s.renamePackageFollow(d, packagePath, newName, out)
			}

			return s.renamePackageThrough(d, packagePath, newName, out)
		},
	)
}
