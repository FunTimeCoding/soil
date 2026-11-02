package go_mod

import (
	"github.com/funtimecoding/soil/pkg/console"
	soil "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/system/run"
)

func Update(
	name string,
	skipProxy bool,
	continueOnError bool,
) {
	console.Format("%s\n", name)
	r := run.New()
	r.Panic = false

	if skipProxy {
		r.Environment(soil.Proxy, soil.Direct)
	}

	r.Start(soil.Go, soil.Get, name)

	if r.Error != nil && IsDeadTag(r.ErrorString) {
		if recovered := recoverDeadTag(
			name,
			r.ErrorString,
			skipProxy,
		); recovered != nil {
			r = recovered
		}
	}

	r.Print()

	if !continueOnError {
		r.PanicOnError()
	}
}
