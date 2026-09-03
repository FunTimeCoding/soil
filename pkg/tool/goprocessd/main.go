package goprocessd

import (
	"github.com/funtimecoding/soil/pkg/argument"
	argumentConstant "github.com/funtimecoding/soil/pkg/argument/constant"
	"github.com/funtimecoding/soil/pkg/instrument"
	"github.com/funtimecoding/soil/pkg/tool/goprocessd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goprocessd/option"
	"github.com/funtimecoding/soil/pkg/web"
)

func Main(
	version string,
	gitHash string,
	buildDate string,
) {
	s := instrument.New(constant.Identity, version)
	defer func() { s.Flush(recover()) }()
	a := argument.NewInstance(constant.Identity)
	a.String(argumentConstant.File, "Procfile", "Path to Procfile")
	a.String("envrc", ".envrc", "Path to .envrc file")
	a.Web()
	a.Parse(version, gitHash, buildDate)
	a.NoPositionals(
		"goprocessd is the daemon and takes no commands - control a running daemon with: goprocess run <command>",
	)
	o := option.New()
	o.ProcfilePath = a.GetString(argumentConstant.File)
	o.EnvrcPath = a.GetString("envrc")
	o.Address = a.Address()
	o.ServiceTokens = web.ServiceTokens()
	o.Version = version
	Run(o, s)
}
