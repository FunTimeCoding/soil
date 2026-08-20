package godashboardd

import (
	"github.com/funtimecoding/soil/pkg/argument"
	argumentConstant "github.com/funtimecoding/soil/pkg/argument/constant"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter"
	"github.com/funtimecoding/soil/pkg/system/environment"
	"github.com/funtimecoding/soil/pkg/tool/godashboardd/board"
	"github.com/funtimecoding/soil/pkg/tool/godashboardd/constant"
	"github.com/funtimecoding/soil/pkg/tool/godashboardd/option"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
)

func Main(
	version string,
	gitHash string,
	buildDate string,
) {
	r := reporter.New(constant.Identity.Name(), version).Start()
	defer func() { r.RecoverFlush(recover()) }()
	a := argument.NewInstance(constant.Identity)
	a.Web()
	a.Database()
	a.String(argumentConstant.Board, constant.BoardFile, constant.BoardUsage)
	a.Parse(version, gitHash, buildDate)
	o := option.New()
	o.Address = a.Address()
	o.Board = board.Load(a.GetString(argumentConstant.Board))
	o.PostgresLocator = a.GetString(argumentConstant.Postgres)
	o.LitePath = a.GetString(argumentConstant.Lite)
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
	o.Version = version
	Run(o, r)
}
