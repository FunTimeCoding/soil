package dump

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/inventory/instance"
	"slices"
)

func (p *Program) Dump(
	x context.Context,
	i *instance.Instance,
	path string,
) error {
	_, e := p.run(
		x,
		i,
		constant.Identity.Name(),
		slices.Concat(
			[]string{constant.DumpProgram, "--format=custom", "--file", path},
			connection(i, i.Database),
		)...,
	)

	return e
}
