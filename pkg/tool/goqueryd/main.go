package goqueryd

import (
	"github.com/funtimecoding/soil/pkg/argument"
	argumentConstant "github.com/funtimecoding/soil/pkg/argument/constant"
	generativeConstant "github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/generative/embed"
	"github.com/funtimecoding/soil/pkg/instrument"
	"github.com/funtimecoding/soil/pkg/system/environment"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/option"
	"github.com/funtimecoding/soil/pkg/web"
)

func Main() {
	s := instrument.New(constant.Identity)
	defer func() { s.Flush(recover()) }()
	a := argument.NewInstance(constant.Identity)
	a.Web()
	a.Metric()
	a.Lite()
	a.Parse()
	o := option.New()
	o.Address = a.Address()
	o.MetricAddress = a.MetricAddress()
	o.ServiceTokens = web.ServiceTokens()
	o.LitePath = a.GetString(argumentConstant.Lite)
	o.EmbedTarget = environment.Required(
		generativeConstant.EmbedTargetEnvironment,
	)
	o.EmbedModel = embed.ModelEnvironment()
	Run(o, s)
}
