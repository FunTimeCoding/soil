package service

import (
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/sink"
	"path/filepath"
)

func (s *Service) MoveSymbols(
	directory string,
	packagePath string,
	symbols []string,
	filePath string,
	targetPackagePath string,
	targetFile string,
	create bool,
	qualifyBackReferences bool,
	dryRun bool,
) (*output.Results, error) {
	absolute := filePath

	if absolute != "" && !filepath.IsAbs(absolute) {
		absolute = filepath.Join(directory, filePath)
	}

	return s.transact(
		directory,
		[]string{packagePath},
		dryRun,
		func(d string, out *sink.Sink) (*output.Results, error) {
			if d != directory {
				return s.moveSymbolsFollow(
					d,
					packagePath,
					symbols,
					absolute,
					targetPackagePath,
					out,
				)
			}

			return s.moveSymbolsThrough(
				d,
				packagePath,
				symbols,
				filePath,
				targetPackagePath,
				targetFile,
				create,
				qualifyBackReferences,
				out,
			)
		},
	)
}
