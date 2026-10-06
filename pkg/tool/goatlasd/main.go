package goatlasd

import (
	"github.com/funtimecoding/soil/pkg/argument"
	argumentConstant "github.com/funtimecoding/soil/pkg/argument/constant"
	"github.com/funtimecoding/soil/pkg/instrument"
	"github.com/funtimecoding/soil/pkg/system/environment"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/option"
	"github.com/funtimecoding/soil/pkg/web"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"time"
)

func Main() {
	s := instrument.New(constant.Identity)
	defer func() { s.Flush(recover()) }()
	a := argument.NewInstance(constant.Identity)
	a.Web()
	a.Metric()
	a.Database()
	a.Duration(argumentConstant.Interval, 5*time.Minute, constant.IntervalUsage)
	a.Duration(
		constant.RetentionArgument,
		7*24*time.Hour,
		constant.RetentionUsage,
	)
	a.Parse()
	o := option.New()
	o.Address = a.Address()
	o.MetricAddress = a.MetricAddress()
	o.ServiceTokens = web.ServiceTokens()
	o.Issuer = environment.Required(webConstant.AuthorizationIssuerEnvironment)
	o.ClientIdentifier = environment.Required(
		webConstant.AuthorizationClientIdentifierEnvironment,
	)
	o.ClientSecret = environment.Required(
		webConstant.AuthorizationClientSecretEnvironment,
	)
	o.EncryptionSecret = environment.Required(
		webConstant.AuthorizationEncryptionSecretEnvironment,
	)
	o.PublicLocator = environment.Required(webConstant.PublicLocatorEnvironment)
	o.LitePath = a.GetString(argumentConstant.Lite)
	o.PostgresLocator = a.GetString(argumentConstant.Postgres)
	o.Interval = a.GetDuration(argumentConstant.Interval)
	o.Retention = a.GetDuration(constant.RetentionArgument)
	Run(o, s)
}
