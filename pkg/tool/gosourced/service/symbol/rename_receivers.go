package symbol

import (
	"context"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/commit"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/sink"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/types/receiver_rename"
)

func (s *Service) RenameReceivers(
	x context.Context,
	directory string,
	packagePath string,
	renames []*receiver_rename.Rename,
	dryRun bool,
) (*output.Results, error) {
	refused := output.NewResultsWithDirectory(directory)

	if len(renames) == 0 {
		return commit.FailValidation(refused, "no receivers given")
	}

	return s.commit.Transact(
		x,
		directory,
		[]string{packagePath},
		dryRun,
		func(d string, out *sink.Sink) (*output.Results, error) {
			return renameReceiversThrough(d, packagePath, renames, out)
		},
	)
}
