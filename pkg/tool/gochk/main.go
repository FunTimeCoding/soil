package gochk

import (
	"github.com/funtimecoding/soil/pkg/argument"
	argumentConstant "github.com/funtimecoding/soil/pkg/argument/constant"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter"
	"github.com/funtimecoding/soil/pkg/tool/gochk/check"
	"github.com/funtimecoding/soil/pkg/tool/gochk/constant"
)

func Main() {
	r := reporter.New(constant.Identity.Name()).Start()
	defer func() { r.RecoverFlush(recover()) }()
	a := argument.NewInstance(constant.Identity)
	a.String(
		argumentConstant.Port,
		"",
		"Port, multiple values separated by comma",
	)
	a.Parse()
	check.Check(a.GetString(argumentConstant.Port))
}
