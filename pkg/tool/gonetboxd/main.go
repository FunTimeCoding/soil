package gonetboxd

import (
	"github.com/funtimecoding/soil/pkg/argument"
	argumentConstant "github.com/funtimecoding/soil/pkg/argument/constant"
	"github.com/funtimecoding/soil/pkg/instrument"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/option"
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
	a.Parse(version, gitHash, buildDate)
	o := option.New()
	o.Address = a.Address()
	o.Version = version
	o.LitePath = a.GetString(argumentConstant.Lite)
	o.PostgresLocator = a.GetString(argumentConstant.Postgres)
	Run(o, s)
}
