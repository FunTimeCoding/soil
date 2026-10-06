package service

import (
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/sink"
)

func (s *Service) ExtractType(
	directory string,
	packagePath string,
	typeName string,
	targetPackagePath string,
	targetFile string,
	create bool,
	dryRun bool,
) (*output.Results, error) {
	return s.transact(
		directory,
		[]string{packagePath},
		dryRun,
		func(d string, out *sink.Sink) (*output.Results, error) {
			if d != directory {
				return s.moveSymbolsFollow(
					d,
					packagePath,
					[]string{typeName},
					"",
					targetPackagePath,
					out,
				)
			}

			return s.extractTypeThrough(
				d,
				packagePath,
				typeName,
				targetPackagePath,
				targetFile,
				create,
				out,
			)
		},
	)
}
