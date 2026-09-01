package go_mod

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/constant"
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
		r.Environment(constant.Proxy, constant.Direct)
	}

	r.Start(constant.Go, constant.Get, name)

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
