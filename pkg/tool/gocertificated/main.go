package gocertificated

import (
	"github.com/funtimecoding/soil/pkg/argument"
	argumentConstant "github.com/funtimecoding/soil/pkg/argument/constant"
	gitlabConstant "github.com/funtimecoding/soil/pkg/gitlab/constant"
	"github.com/funtimecoding/soil/pkg/instrument"
	"github.com/funtimecoding/soil/pkg/system/environment"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/constant"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/option"
)

func Main() {
	s := instrument.New(constant.Identity)
	defer func() { s.Flush(recover()) }()
	a := argument.NewInstance(constant.Identity)
	a.Web()
	a.Database()
	a.Parse()
	o := option.New()
	o.Read(a)
	o.PostgresLocator = a.GetString(argumentConstant.Postgres)
	o.LitePath = a.GetString(argumentConstant.Lite)
	o.Project = environment.Required(gitlabConstant.ProjectNameEnvironment)
	o.Branch = constant.PublishBranch
	o.AuthorityDirectory = environment.Required(
		constant.AuthorityDirectoryEnvironment,
	)
	o.SecretAuthority = environment.Required(
		constant.SecretAuthorityEnvironment,
	)
	o.SecretPath = environment.Required(constant.SecretPathEnvironment)
	Run(o, s)
}
