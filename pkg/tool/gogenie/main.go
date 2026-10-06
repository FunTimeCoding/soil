package gogenie

import (
	"github.com/funtimecoding/soil/pkg/argument"
	argumentConstant "github.com/funtimecoding/soil/pkg/argument/constant"
	"github.com/funtimecoding/soil/pkg/atlassian/opsgenie/check/alert"
	"github.com/funtimecoding/soil/pkg/atlassian/opsgenie/check/alert/option"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter"
	"github.com/funtimecoding/soil/pkg/tool/gogenie/constant"
)

func Main() {
	r := reporter.New(constant.Identity.Name()).Start()
	defer func() { r.RecoverFlush(recover()) }()
	a := argument.NewInstance(constant.Identity)
	a.Boolean(
		argumentConstant.Copyable,
		false,
		"Disable OSC8 links and add a copyable link instead",
	)
	a.Boolean(argumentConstant.Notation, false, "JSON output")
	a.Boolean(argumentConstant.All, false, "Include filtered in output")
	a.Parse()
	o := option.New()
	o.Copyable = a.GetBoolean(argumentConstant.Copyable)
	o.Notation = a.GetBoolean(argumentConstant.Notation)
	o.All = a.GetBoolean(argumentConstant.All)
	alert.Check(o)
}
