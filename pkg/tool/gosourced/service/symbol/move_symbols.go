package symbol

import (
	"context"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/sink"
	"path/filepath"
)

func (s *Service) MoveSymbols(
	x context.Context,
	directory string,
	packagePath string,
	symbols []string,
	filePath string,
	targetPackagePath string,
	targetFile string,
	create bool,
	dryRun bool,
) (*output.Results, error) {
	absolute := filePath

	if absolute != "" && !filepath.IsAbs(absolute) {
		absolute = filepath.Join(directory, filePath)
	}

	return s.commit.Transact(
		x,
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
				out,
			)
		},
	)
}
