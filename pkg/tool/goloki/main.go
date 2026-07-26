package goloki

import (
	"github.com/funtimecoding/soil/pkg/argument"
	argumentConstant "github.com/funtimecoding/soil/pkg/argument/constant"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter"
	"github.com/funtimecoding/soil/pkg/prometheus/check/loki"
	"github.com/funtimecoding/soil/pkg/prometheus/check/loki/option"
	lokiConstant "github.com/funtimecoding/soil/pkg/prometheus/loki/constant"
	"github.com/funtimecoding/soil/pkg/system/environment"
	"github.com/funtimecoding/soil/pkg/tool/goloki/constant"
	"time"
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
	a.Duration(argumentConstant.Since, time.Hour, "Time range to query")
	a.String(argumentConstant.Route, "", "Filter by HTTP route")
	a.String(argumentConstant.Message, "", "Filter by message field")
	a.BooleanShort(argumentConstant.Body, "b", false, "Output only the body field")
	a.IntegerShort(
		argumentConstant.Limit,
		"n",
		10,
		"Maximum number of log entries",
	)
	a.Parse(version, gitHash, buildDate)
	o := option.New()
	o.Namespace = a.Argument(0)
	o.Since = a.GetDuration(argumentConstant.Since)
	o.Route = a.GetString(argumentConstant.Route)
	o.Message = a.GetString(argumentConstant.Message)
	o.BodyOnly = a.GetBoolean(argumentConstant.Body)
	o.Copyable = a.GetBoolean(argumentConstant.Copyable)
	o.Limit = a.GetInteger(argumentConstant.Limit)
	o.Namespaces = environment.Slice(lokiConstant.NamespaceEnvironment)
	o.Exclude = environment.Slice(lokiConstant.ExcludeEnvironment)
	loki.Check(o)
}
