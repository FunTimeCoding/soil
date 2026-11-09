package mock_dump

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/inventory/instance"
)

func (p *Program) Restore(
	_ context.Context,
	_ *instance.Instance,
	database string,
	path string,
	clean bool,
) error {
	p.Restored = append(
		p.Restored,
		&Restoration{Database: database, Path: path, Clean: clean},
	)

	return p.Failure
}
