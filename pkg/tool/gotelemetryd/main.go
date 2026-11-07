package gotelemetryd

import (
	"github.com/funtimecoding/soil/pkg/argument"
	argumentConstant "github.com/funtimecoding/soil/pkg/argument/constant"
	"github.com/funtimecoding/soil/pkg/instrument"
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/option"
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
	Run(o, s)
}
