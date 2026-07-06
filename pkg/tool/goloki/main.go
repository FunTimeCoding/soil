package goloki

import (
	"github.com/funtimecoding/soil/pkg/argument"
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
		argument.Copyable,
		false,
		"Disable OSC8 links and add a copyable link instead",
	)
	a.Duration(argument.Since, time.Hour, "Time range to query")
	a.String(argument.Route, "", "Filter by HTTP route")
	a.String(argument.Message, "", "Filter by message field")
	a.BooleanShort(argument.Body, "b", false, "Output only the body field")
	a.IntegerShort(argument.Limit, "n", 10, "Maximum number of log entries")
	a.Parse(version, gitHash, buildDate)
	o := option.New()
	o.Namespace = a.Argument(0)
	o.Since = a.GetDuration(argument.Since)
	o.Route = a.GetString(argument.Route)
	o.Message = a.GetString(argument.Message)
	o.BodyOnly = a.GetBoolean(argument.Body)
	o.Copyable = a.GetBoolean(argument.Copyable)
	o.Limit = a.GetInteger(argument.Limit)
	o.Namespaces = environment.Slice(lokiConstant.NamespaceEnvironment)
	o.Exclude = environment.Slice(lokiConstant.ExcludeEnvironment)
	loki.Check(o)
}
