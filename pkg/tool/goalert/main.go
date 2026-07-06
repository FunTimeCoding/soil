package goalert

import (
	"github.com/funtimecoding/soil/pkg/argument"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter"
	"github.com/funtimecoding/soil/pkg/prometheus/check/alert"
	"github.com/funtimecoding/soil/pkg/prometheus/check/alert/option"
	"github.com/funtimecoding/soil/pkg/tool/goalert/constant"
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
	a.Boolean(argument.Critical, false, "Critical severity only")
	a.Boolean(argument.Warning, false, "Warning severity only")
	a.Boolean(argument.Extended, false, "Extended output")
	a.Boolean(argument.Suppressed, false, "Include suppressed")
	a.Boolean(argument.Rules, false, "Print rules")
	a.Boolean(argument.Firing, false, "Print firing rules")
	a.Boolean(argument.Fingerprint, false, "Fingerprint column")
	a.Parse(version, gitHash, buildDate)
	o := option.New()
	o.Notation = a.GetBoolean(argument.Notation)
	o.All = a.GetBoolean(argument.All)
	o.Critical = a.GetBoolean(argument.Critical)
	o.Warning = a.GetBoolean(argument.Warning)
	o.Extended = a.GetBoolean(argument.Extended)
	o.Suppressed = a.GetBoolean(argument.Suppressed)
	o.Rules = a.GetBoolean(argument.Rules)
	o.Firing = a.GetBoolean(argument.Firing)
	o.Fingerprint = a.GetBoolean(argument.Fingerprint)
	o.Copyable = a.GetBoolean(argument.Copyable)
	alert.Check(o)
}
