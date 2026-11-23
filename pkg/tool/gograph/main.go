package gograph

import (
	"github.com/funtimecoding/soil/pkg/argument"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/tool/gograph/constant"
)

func Main() {
	r := reporter.New(constant.Identity.Name()).Start()
	defer func() { r.RecoverFlush(recover()) }()
	a := argument.NewInstance(constant.Identity)
	a.Parse()

	if c := a.RequiredPositional(0, "COMMAND"); c != constant.Login {
		system.Exitf(1, "unknown command: %s\n", c)
	}

	Run()
}
