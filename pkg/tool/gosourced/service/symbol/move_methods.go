package symbol

import (
	"context"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/sink"
)

func (s *Service) MoveMethods(
	x context.Context,
	directory string,
	packagePath string,
	typeName string,
	methods []string,
	targetPackagePath string,
	targetType string,
	targetFile string,
	dryRun bool,
) (*output.Results, error) {
	return s.commit.Transact(
		x,
		directory,
		[]string{packagePath},
		dryRun,
		func(d string, out *sink.Sink) (*output.Results, error) {
			if d != directory {
				return s.moveMethodsFollow(d, packagePath, typeName, methods)
			}

			return s.moveMethodsThrough(
				d,
				packagePath,
				typeName,
				methods,
				targetPackagePath,
				targetType,
				targetFile,
				out,
			)
		},
	)
}
