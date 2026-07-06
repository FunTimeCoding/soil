package gobrew

import (
	"github.com/funtimecoding/soil/pkg/argument"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter"
	"github.com/funtimecoding/soil/pkg/system/macos/check/brew/outdated"
	"github.com/funtimecoding/soil/pkg/system/macos/check/brew/outdated/option"
	"github.com/funtimecoding/soil/pkg/tool/gobrew/constant"
)

func Main(
	version string,
	gitHash string,
	buildDate string,
) {
	r := reporter.New(constant.Identity.Name(), version).Start()
	defer func() { r.RecoverFlush(recover()) }()
	a := argument.NewInstance(constant.Identity)
	a.Boolean(
		argument.Copyable,
		false,
		"Disable OSC8 links and add a copyable link instead",
	)
	a.Boolean(argument.Notation, false, "JSON output")
	a.Boolean(argument.All, false, "Include filtered in output")
	a.Parse(version, gitHash, buildDate)
	o := option.New()
	o.Copyable = a.GetBoolean(argument.Copyable)
	o.Notation = a.GetBoolean(argument.Notation)
	o.All = a.GetBoolean(argument.All)
	outdated.Check(o)
}
