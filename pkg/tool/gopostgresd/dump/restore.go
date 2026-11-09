package dump

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/inventory/instance"
	"slices"
)

func (p *Program) Restore(
	x context.Context,
	i *instance.Instance,
	database string,
	path string,
	clean bool,
) error {
	a := []string{constant.RestoreProgram, "--single-transaction"}

	if clean {
		a = append(a, "--clean", "--if-exists")
	}

	_, e := p.run(
		x,
		i,
		constant.RestoreApplicationName,
		slices.Concat(a, connection(i, database), []string{path})...,
	)

	return e
}
