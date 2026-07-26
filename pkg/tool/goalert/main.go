package goalert

import (
	"github.com/funtimecoding/soil/pkg/argument"
	argumentConstant "github.com/funtimecoding/soil/pkg/argument/constant"
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
		argumentConstant.Copyable,
		false,
		"Disable OSC8 links and add a copyable link instead",
	)
	a.Boolean(argumentConstant.Notation, false, "JSON output")
	a.Boolean(argumentConstant.All, false, "Include filtered in output")
	a.Boolean(argumentConstant.Critical, false, "Critical severity only")
	a.Boolean(argumentConstant.Warning, false, "Warning severity only")
	a.Boolean(argumentConstant.Extended, false, "Extended output")
	a.Boolean(argumentConstant.Suppressed, false, "Include suppressed")
	a.Boolean(argumentConstant.Rules, false, "Print rules")
	a.Boolean(argumentConstant.Firing, false, "Print firing rules")
	a.Boolean(argumentConstant.Fingerprint, false, "Fingerprint column")
	a.Parse(version, gitHash, buildDate)
	o := option.New()
	o.Notation = a.GetBoolean(argumentConstant.Notation)
	o.All = a.GetBoolean(argumentConstant.All)
	o.Critical = a.GetBoolean(argumentConstant.Critical)
	o.Warning = a.GetBoolean(argumentConstant.Warning)
	o.Extended = a.GetBoolean(argumentConstant.Extended)
	o.Suppressed = a.GetBoolean(argumentConstant.Suppressed)
	o.Rules = a.GetBoolean(argumentConstant.Rules)
	o.Firing = a.GetBoolean(argumentConstant.Firing)
	o.Fingerprint = a.GetBoolean(argumentConstant.Fingerprint)
	o.Copyable = a.GetBoolean(argumentConstant.Copyable)
	alert.Check(o)
}
