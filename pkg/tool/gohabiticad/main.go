package gohabiticad

import (
	"github.com/funtimecoding/soil/pkg/argument"
	habiticaConstant "github.com/funtimecoding/soil/pkg/habitica/constant"
	"github.com/funtimecoding/soil/pkg/instrument"
	"github.com/funtimecoding/soil/pkg/system/environment"
	"github.com/funtimecoding/soil/pkg/tool/gohabiticad/constant"
	"github.com/funtimecoding/soil/pkg/tool/gohabiticad/option"
)

func Main() {
	s := instrument.New(constant.Identity)
	defer func() { s.Flush(recover()) }()
	a := argument.NewInstance(constant.Identity)
	a.Web()
	a.Parse()
	o := option.New()
	o.Read(a)
	o.Host = environment.Required(habiticaConstant.HostEnvironment)
	Run(o, s)
}
