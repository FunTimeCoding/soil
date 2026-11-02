package packages

import (
	"context"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/sink"
)

func (s *Service) RenamePackageClause(
	x context.Context,
	directory string,
	packagePath string,
	newName string,
	dryRun bool,
) (*output.Results, error) {
	return s.commit.Transact(
		x,
		directory,
		[]string{packagePath},
		dryRun,
		func(d string, out *sink.Sink) (*output.Results, error) {
			if d != directory {
				return s.renamePackageClauseFollow(d, packagePath, newName, out)
			}

			return s.renamePackageClauseThrough(d, packagePath, newName, out)
		},
	)
}
