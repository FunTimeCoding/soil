package gochk

import (
	"github.com/funtimecoding/soil/pkg/argument"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter"
	"github.com/funtimecoding/soil/pkg/tool/gochk/check"
	"github.com/funtimecoding/soil/pkg/tool/gochk/constant"
)

func Main(
	version string,
	gitHash string,
	buildDate string,
) {
	r := reporter.New(constant.Identity.Name(), version).Start()
	defer func() { r.RecoverFlush(recover()) }()
	a := argument.NewInstance(constant.Identity)
	a.String(
		argument.Port,
		"",
		"Port, multiple values separated by comma",
	)
	a.Parse(version, gitHash, buildDate)
	check.Check(a.GetString(argument.Port))
}
