package gosilence

import (
	"github.com/funtimecoding/soil/pkg/argument"
	argumentConstant "github.com/funtimecoding/soil/pkg/argument/constant"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/check/silence"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/check/silence/option"
	"github.com/funtimecoding/soil/pkg/tool/gosilence/constant"
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
		argumentConstant.Copyable,
		false,
		"Disable OSC8 links and add a copyable link instead",
	)
	a.Boolean(argumentConstant.Notation, false, "JSON output")
	a.Boolean(argumentConstant.All, false, "Include filtered in output")
	a.String(argumentConstant.Set, "", "Name, creates or updates")
	a.String(argumentConstant.Duration, "", "Duration, default 10m")
	a.Parse(version, gitHash, buildDate)
	o := option.New()
	o.Notation = a.GetBoolean(argumentConstant.Notation)
	o.All = a.GetBoolean(argumentConstant.All)
	o.Set = a.GetString(argumentConstant.Set)
	o.Copyable = a.GetBoolean(argumentConstant.Copyable)
	silence.Check(o)
}
