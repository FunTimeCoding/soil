package gofirefoxd

import (
	"github.com/funtimecoding/soil/pkg/argument"
	"github.com/funtimecoding/soil/pkg/instrument"
	"github.com/funtimecoding/soil/pkg/tool/gofirefoxd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gofirefoxd/option"
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
	a.Integer(
		constant.BridgePortFlag,
		6125,
		"WebSocket bridge port for extension",
	)
	a.Parse(version, gitHash, buildDate)
	o := option.New()
	o.Address = a.Address()
	o.BridgePort = a.RequiredInteger(constant.BridgePortFlag)
	o.Version = version
	Run(o, s)
}
