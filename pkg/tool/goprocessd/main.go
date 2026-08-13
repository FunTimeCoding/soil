package goprocessd

import (
	"github.com/funtimecoding/soil/pkg/argument"
	argumentConstant "github.com/funtimecoding/soil/pkg/argument/constant"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter"
	"github.com/funtimecoding/soil/pkg/tool/goprocessd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goprocessd/option"
)

func Main(
	version string,
	gitHash string,
	buildDate string,
) {
	r := reporter.New(constant.Identity.Name(), version)
	r.Start()
	defer func() { r.RecoverFlush(recover()) }()
	a := argument.NewInstance(constant.Identity)
	a.String(argumentConstant.File, "Procfile", "Path to Procfile")
	a.String("envrc", ".envrc", "Path to .envrc file")
	a.Parse(version, gitHash, buildDate)
	a.NoPositionals(
		"goprocessd is the daemon and takes no commands - control a running daemon with: goprocess run <command>",
	)
	o := option.New()
	o.ProcfilePath = a.GetString(argumentConstant.File)
	o.EnvrcPath = a.GetString("envrc")
	Run(o)
}
