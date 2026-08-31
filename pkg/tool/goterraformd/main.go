package goterraformd

import (
	"github.com/funtimecoding/soil/pkg/argument"
	argumentConstant "github.com/funtimecoding/soil/pkg/argument/constant"
	"github.com/funtimecoding/soil/pkg/instrument"
	"github.com/funtimecoding/soil/pkg/strings/split"
	"github.com/funtimecoding/soil/pkg/system/environment"
	"github.com/funtimecoding/soil/pkg/tool/goterraformd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goterraformd/option"
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
	a.Parse(version, gitHash, buildDate)
	o := option.New()
	o.Address = a.Address()
	o.Version = version
	o.Repository = a.Required(argumentConstant.Repository)
	o.ClonePath = a.Required(argumentConstant.ClonePath)
	o.TerraformPath = environment.Required(constant.TerraformPathEnvironment)
	o.PostgresLocator = a.GetString(argumentConstant.Postgres)
	o.LitePath = a.GetString(argumentConstant.Lite)
	o.StateNamespace = environment.Fallback(
		constant.StateNamespaceEnvironment,
		constant.StateNamespace,
	)
	o.StateLease = environment.Fallback(
		constant.StateLeaseEnvironment,
		constant.StateLease,
	)

	if v := environment.Fallback(constant.DownstreamEnvironment, ""); v != "" {
		o.Downstream = split.Comma(v)
	}

	Run(o, s)
}
