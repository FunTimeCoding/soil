package service

import (
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/sink"
)

func (s *Service) Rename(
	directory string,
	packagePath string,
	oldName string,
	newName string,
	receiver string,
	dryRun bool,
) (*output.Results, error) {
	return s.transact(
		directory,
		[]string{packagePath},
		dryRun,
		func(d string, out *sink.Sink) (*output.Results, error) {
			return s.renameThrough(
				d,
				packagePath,
				oldName,
				newName,
				receiver,
				out,
			)
		},
	)
}
