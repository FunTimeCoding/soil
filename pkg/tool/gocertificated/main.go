package gocertificated

import (
	"github.com/funtimecoding/soil/pkg/argument"
	argumentConstant "github.com/funtimecoding/soil/pkg/argument/constant"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter"
	gitlabConstant "github.com/funtimecoding/soil/pkg/gitlab/constant"
	"github.com/funtimecoding/soil/pkg/system/environment"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/constant"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/option"
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
	a.String(
		constant.ProjectArgument,
		environment.Fallback(gitlabConstant.ProjectEnvironment, ""),
		"Repository the authority publishes to",
	)
	a.String(
		constant.BranchArgument,
		constant.PublishBranch,
		"Branch the authority commits to",
	)
	a.String(
		constant.SecretAuthorityArgument,
		environment.Fallback(constant.SecretAuthorityEnvironment, ""),
		"Authority whose material is delivered as a secret manifest",
	)
	a.String(
		constant.SecretPathArgument,
		environment.Fallback(constant.SecretPathEnvironment, ""),
		"Repository path of the delivered secret manifest",
	)
	a.Parse(version, gitHash, buildDate)
	o := option.New()
	o.Address = a.Address()
	o.PostgresLocator = a.GetString(argumentConstant.Postgres)
	o.LitePath = a.GetString(argumentConstant.Lite)
	o.Project = a.GetString(constant.ProjectArgument)
	o.Branch = a.GetString(constant.BranchArgument)
	o.SecretAuthority = a.GetString(constant.SecretAuthorityArgument)
	o.SecretPath = a.GetString(constant.SecretPathArgument)
	o.Version = version
	Run(o, r)
}
