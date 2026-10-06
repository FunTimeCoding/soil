package gogitlab

import (
	"github.com/funtimecoding/soil/pkg/argument"
	argumentConstant "github.com/funtimecoding/soil/pkg/argument/constant"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter"
	"github.com/funtimecoding/soil/pkg/gitlab/check/job"
	"github.com/funtimecoding/soil/pkg/gitlab/check/job/option"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/constant"
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
	a.Boolean(argumentConstant.Verbose, false, "Verbose output")
	a.Parse()
	o := option.New()
	o.Notation = a.GetBoolean(argumentConstant.Notation)
	o.All = a.GetBoolean(argumentConstant.All)
	o.Verbose = a.GetBoolean(argumentConstant.Verbose)
	o.Copyable = a.GetBoolean(argumentConstant.Copyable)
	job.Check(o)
}
