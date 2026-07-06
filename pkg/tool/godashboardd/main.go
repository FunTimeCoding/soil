package godashboardd

import (
	"github.com/funtimecoding/soil/pkg/argument"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter"
	"github.com/funtimecoding/soil/pkg/relational/postgres"
	"github.com/funtimecoding/soil/pkg/system/environment"
	"github.com/funtimecoding/soil/pkg/tool/godashboardd/board"
	"github.com/funtimecoding/soil/pkg/tool/godashboardd/constant"
	"github.com/funtimecoding/soil/pkg/tool/godashboardd/option"
	web "github.com/funtimecoding/soil/pkg/web/constant"
)

func Main(
	version string,
	gitHash string,
	buildDate string,
) {
	r := reporter.New(constant.Identity.Name(), version).Start()
	defer func() { r.RecoverFlush(recover()) }()
	a := argument.NewInstance(constant.Identity)
	a.Integer(argument.Port, web.ListenPort, web.PortUsage)
	a.String(argument.Board, constant.BoardFile, constant.BoardUsage)
	a.String(argument.Path, "", "SQLite database path")
	a.Parse(version, gitHash, buildDate)
	o := option.New()
	o.Port = a.RequiredInteger(argument.Port)
	o.Board = board.Load(a.GetString(argument.Board))
	o.PostgresLocator = environment.Optional(postgres.LocatorEnvironment)
	o.LitePath = a.GetString(argument.Path)
	o.Issuer = environment.Optional(constant.IssuerEnvironment)
	o.ClientIdentifier = environment.Optional(
		constant.ClientIdentifierEnvironment,
	)
	o.ClientSecret = environment.Optional(
		constant.ClientSecretEnvironment,
	)
	o.EncryptionSecret = environment.Optional(
		constant.EncryptionSecretEnvironment,
	)
	o.PublicLocator = environment.Optional(
		constant.PublicLocatorEnvironment,
	)
	o.Version = version
	Run(o, r)
}
