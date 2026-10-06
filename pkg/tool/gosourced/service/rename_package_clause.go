package service

import (
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/sink"
)

func (s *Service) RenamePackageClause(
	directory string,
	packagePath string,
	newName string,
	dryRun bool,
) (*output.Results, error) {
	return s.transact(
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
