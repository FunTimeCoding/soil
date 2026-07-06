package gokevt

import (
	"github.com/funtimecoding/soil/pkg/argument"
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
	a.Boolean(argument.Notation, false, "JSON output")
	a.Boolean(argument.All, false, "Include filtered in output")
	a.Boolean(argument.Clean, false, "Delete events older than 7 days")
	a.Parse(version, gitHash, buildDate)
	o := option.New()
	o.Notation = a.GetBoolean(argument.Notation)
	o.All = a.GetBoolean(argument.All)
	o.Clean = a.GetBoolean(argument.Clean)
	event.Print(o)
}
