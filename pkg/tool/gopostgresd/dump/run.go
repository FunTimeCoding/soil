package dump

import (
	"context"
	"github.com/funtimecoding/soil/pkg/system/run"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/inventory/instance"
	"slices"
)

func (p *Program) run(
	x context.Context,
	i *instance.Instance,
	application string,
	a ...string,
) (string, error) {
	r := run.New().NoPanic().WithContext(x)

	if i != nil {
		r.Environment(constant.PasswordEnvironment, i.Password)
	}

	r.Environment(constant.ApplicationNameEnvironment, application)
	result := r.Start(slices.Concat(p.program, a)...)

	return result, r.Error
}
