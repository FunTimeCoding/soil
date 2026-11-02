package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/go_mod/constant"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/sink"
)

func (s *Service) MovePackage(
	x context.Context,
	directory string,
	packagePath string,
	targetPackagePath string,
	dryRun bool,
) (*output.Results, error) {
	return s.transact(
		x,
		directory,
		[]string{join.Empty(packagePath, constant.RecursivePattern)},
		dryRun,
		func(d string, out *sink.Sink) (*output.Results, error) {
			if d != directory {
				return s.movePackageFollow(
					d,
					packagePath,
					targetPackagePath,
					out,
				)
			}

			return s.movePackageThrough(d, packagePath, targetPackagePath, out)
		},
	)
}
