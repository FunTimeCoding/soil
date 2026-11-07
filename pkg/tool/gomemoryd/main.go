package gomemoryd

import (
	"github.com/funtimecoding/soil/pkg/argument"
	argumentConstant "github.com/funtimecoding/soil/pkg/argument/constant"
	"github.com/funtimecoding/soil/pkg/instrument"
	"github.com/funtimecoding/soil/pkg/system/environment"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/option"
)

func Main() {
	s := instrument.New(constant.Identity)
	defer func() { s.Flush(recover()) }()
	a := argument.NewInstance(constant.Identity)
	a.Web()
	a.Lite()
	a.Parse()
	o := option.New()
	o.Read(a)
	o.LitePath = a.GetString(argumentConstant.Lite)
	o.HiddenTag = environment.Fallback(constant.HiddenTagEnvironment, "")
	o.ReferenceRoot = environment.Fallback(
		constant.ReferenceRootEnvironment,
		"",
	)
	Run(o, s)
}
