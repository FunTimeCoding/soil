package goraidd

import (
	"github.com/funtimecoding/soil/pkg/argument"
	argumentConstant "github.com/funtimecoding/soil/pkg/argument/constant"
	"github.com/funtimecoding/soil/pkg/instrument"
	"github.com/funtimecoding/soil/pkg/system/environment"
	"github.com/funtimecoding/soil/pkg/tool/goraidd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goraidd/option"
	"github.com/funtimecoding/soil/pkg/web"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
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
	o.PostgresLocator = a.GetString(argumentConstant.Postgres)
	o.LitePath = a.GetString(argumentConstant.Lite)
	o.LogCachePath = "/srv/arcdps-config"
	o.ElitePath = "/srv/elite-insights"
	o.OutputPath = "/srv/gw2-report"
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
	o.ServiceTokens = web.ServiceTokens()
	Run(o, s)
}
