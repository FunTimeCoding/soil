package symbol

import (
	"context"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/symbol/identifier"
)

func (s *Service) ChangeVisibility(
	x context.Context,
	directory string,
	symbol string,
	packagePath string,
	receiver string,
	dryRun bool,
) (*output.Results, error) {
	return s.Rename(
		x,
		directory,
		packagePath,
		symbol,
		identifier.FlipName(symbol),
		receiver,
		dryRun,
	)
}
