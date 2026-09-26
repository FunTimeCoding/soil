package goagentd

import (
	"github.com/funtimecoding/soil/pkg/argument"
	soil "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/instrument"
	"github.com/funtimecoding/soil/pkg/tool/goagentd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goagentd/option"
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
	a.Web()
	a.String(
		"workspace",
		soil.CurrentDirectory,
		"Workspace directory for Claude Code",
	)
	a.Parse(version, gitHash, buildDate)
	o := option.New()
	o.Address = a.Address()
	o.Workspace = a.GetString("workspace")
	o.ServiceTokens = web.ServiceTokens()
	Run(o, s)
}
