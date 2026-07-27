package gokevt

import (
	"github.com/funtimecoding/soil/pkg/argument"
	argumentConstant "github.com/funtimecoding/soil/pkg/argument/constant"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter"
	"github.com/funtimecoding/soil/pkg/kubernetes/check/event"
	"github.com/funtimecoding/soil/pkg/kubernetes/check/event/option"
	"github.com/funtimecoding/soil/pkg/tool/gokevt/constant"
)

func Main(
	version string,
	gitHash string,
	buildDate string,
) {
	r := reporter.New(constant.Identity.Name(), version).Start()
	defer func() { r.RecoverFlush(recover()) }()
	a := argument.NewInstance(constant.Identity)
	a.Boolean(argumentConstant.Notation, false, "JSON output")
	a.Boolean(argumentConstant.All, false, "Include filtered in output")
	a.Boolean(
		argumentConstant.Clean,
		false,
		"Delete events older than 7 days",
	)
	a.Parse(version, gitHash, buildDate)
	o := option.New()
	o.Notation = a.GetBoolean(argumentConstant.Notation)
	o.All = a.GetBoolean(argumentConstant.All)
	o.Clean = a.GetBoolean(argumentConstant.Clean)
	event.Print(o)
}
