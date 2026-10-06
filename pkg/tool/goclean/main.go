package goclean

import (
	"github.com/funtimecoding/soil/pkg/argument"
	argumentConstant "github.com/funtimecoding/soil/pkg/argument/constant"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter"
	gitlab "github.com/funtimecoding/soil/pkg/gitlab/constant"
	"github.com/funtimecoding/soil/pkg/system/environment"
	"github.com/funtimecoding/soil/pkg/tool/goclean/clean"
	"github.com/funtimecoding/soil/pkg/tool/goclean/clean/option"
	"github.com/funtimecoding/soil/pkg/tool/goclean/constant"
)

func Main() {
	r := reporter.New(constant.Identity.Name()).Start()
	defer func() { r.RecoverFlush(recover()) }()
	a := argument.NewInstance(constant.Identity)
	a.Boolean(argumentConstant.Verbose, false, "Verbose output")
	a.Boolean(
		argumentConstant.All,
		false,
		"Delete running and queued pipelines too",
	)
	a.Parse()
	o := option.New()
	o.GitLabHost = environment.Required(gitlab.HostEnvironment)
	o.Verbose = a.GetBoolean(argumentConstant.Verbose)
	o.All = a.GetBoolean(argumentConstant.All)
	clean.Run(o)
}
