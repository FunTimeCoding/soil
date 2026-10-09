package gosentryd

import (
	"github.com/funtimecoding/soil/pkg/argument"
	errors "github.com/funtimecoding/soil/pkg/errors/constant"
	"github.com/funtimecoding/soil/pkg/instrument"
	"github.com/funtimecoding/soil/pkg/system/environment"
	"github.com/funtimecoding/soil/pkg/tool/gosentryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosentryd/option"
	"github.com/funtimecoding/soil/pkg/web"
)

func Main() {
	s := instrument.New(constant.Identity)
	defer func() { s.Flush(recover()) }()
	a := argument.NewInstance(constant.Identity)
	a.Web()
	a.Parse()
	o := option.New()
	o.Address = a.Address()
	o.ServiceTokens = web.ServiceTokens()
	o.Organization = environment.Required(errors.OrganizationEnvironment)
	o.Host = environment.Required(errors.HostEnvironment)
	Run(o, s)
}
