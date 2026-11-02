package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/go_mod/constant"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/sink"
)

func (s *Service) RenamePackage(
	x context.Context,
	directory string,
	packagePath string,
	newName string,
	dryRun bool,
) (*output.Results, error) {
	return s.transact(
		x,
		directory,
		[]string{join.Empty(packagePath, constant.RecursivePattern)},
		dryRun,
		func(d string, out *sink.Sink) (*output.Results, error) {
			if d != directory {
				return s.renamePackageFollow(d, packagePath, newName, out)
			}

			return s.renamePackageThrough(d, packagePath, newName, out)
		},
	)
}
