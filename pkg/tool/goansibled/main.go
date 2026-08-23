package goansibled

import (
	"github.com/funtimecoding/soil/pkg/argument"
	argumentConstant "github.com/funtimecoding/soil/pkg/argument/constant"
	"github.com/funtimecoding/soil/pkg/instrument"
	"github.com/funtimecoding/soil/pkg/system/environment"
	"github.com/funtimecoding/soil/pkg/tool/goansibled/constant"
	"github.com/funtimecoding/soil/pkg/tool/goansibled/option"
)

func Main(
	version string,
	gitHash string,
	buildDate string,
) {
	s := instrument.New(constant.Identity, version)
	defer func() { s.Flush(recover()) }()
	a := argument.NewInstance(constant.Identity)
	a.Web()
	a.Database()
	a.String(argumentConstant.Repository, "", "Git repository URL")
	a.String(argumentConstant.ClonePath, "", "Local repository path")
	a.String(
		argumentConstant.Playbook,
		"",
		"Comma-separated playbook paths relative to ansible path",
	)
	a.Parse(version, gitHash, buildDate)
	o := option.New()
	o.Address = a.Address()
	o.Version = version
	o.Repository = a.Required(argumentConstant.Repository)
	o.ClonePath = a.Required(argumentConstant.ClonePath)
	o.AnsiblePath = environment.Required(constant.AnsiblePathEnvironment)
	o.Playbook = a.Slice(argumentConstant.Playbook)
	o.PostgresLocator = a.GetString(argumentConstant.Postgres)
	o.LitePath = a.GetString(argumentConstant.Lite)
	Run(o, s)
}
