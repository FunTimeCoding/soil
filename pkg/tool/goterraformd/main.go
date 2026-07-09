package goterraformd

import (
	"github.com/funtimecoding/soil/pkg/argument"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter"
	"github.com/funtimecoding/soil/pkg/system/environment"
	"github.com/funtimecoding/soil/pkg/tool/goterraformd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goterraformd/option"
)

func Main(
	version string,
	gitHash string,
	buildDate string,
) {
	r := reporter.New(constant.Identity.Name(), version).Start()
	defer func() { r.RecoverFlush(recover()) }()
	a := argument.NewInstance(constant.Identity)
	a.Web()
	a.Database()
	a.String(argument.Repository, "", "Git repository URL")
	a.String(argument.ClonePath, "", "Local repository path")
	a.Parse(version, gitHash, buildDate)
	o := option.New()
	o.Address = a.Address()
	o.Version = version
	o.Repository = a.Required(argument.Repository)
	o.ClonePath = a.Required(argument.ClonePath)
	o.TerraformPath = environment.Required(constant.TerraformPathEnvironment)
	o.PostgresLocator = a.GetString(argument.Postgres)
	o.LitePath = a.GetString(argument.Lite)
	Run(o, r)
}
