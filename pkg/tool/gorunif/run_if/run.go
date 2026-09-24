package run_if

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/git/changed"
	"github.com/funtimecoding/soil/pkg/system/constant"
	"github.com/funtimecoding/soil/pkg/system/run"
	"github.com/funtimecoding/soil/pkg/tool/gorunif/run_if/option"
)

func Run(o *option.RunIf) bool {
	var r *changed.Range

	if o.Base == "" {
		r = changed.NewPush(o.Directory)
	} else {
		r = changed.New(o.Base, o.Head)
	}

	if !r.All && !Matches(o, r.Files(o.Directory)) {
		return false
	}

	e := run.New()
	e.Directory = o.Directory
	e.Execute(constant.Shell, constant.ShellCommand, o.Execute)

	if o.Verbose {
		console.Format("Ran: %s\n", o.Execute)
	}

	return true
}
