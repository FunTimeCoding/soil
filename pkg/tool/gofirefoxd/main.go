package gofirefoxd

import (
	"github.com/funtimecoding/soil/pkg/argument"
	"github.com/funtimecoding/soil/pkg/instrument"
	"github.com/funtimecoding/soil/pkg/tool/gofirefoxd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gofirefoxd/option"
	"github.com/funtimecoding/soil/pkg/web"
)

func Main() {
	s := instrument.New(constant.Identity)
	defer func() { s.Flush(recover()) }()
	a := argument.NewInstance(constant.Identity)
	a.Web()
	a.Integer(
		constant.BridgePortFlag,
		6125,
		"WebSocket bridge port for extension",
	)
	a.Parse()
	o := option.New()
	o.Address = a.Address()
	o.ServiceTokens = web.ServiceTokens()
	o.BridgePort = a.RequiredInteger(constant.BridgePortFlag)
	Run(o, s)
}
