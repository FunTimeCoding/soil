package godashboardd

import (
	"github.com/funtimecoding/soil/pkg/argument"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter"
	"github.com/funtimecoding/soil/pkg/system/environment"
	"github.com/funtimecoding/soil/pkg/tool/godashboardd/board"
	"github.com/funtimecoding/soil/pkg/tool/godashboardd/constant"
	"github.com/funtimecoding/soil/pkg/tool/godashboardd/option"
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
	a.String(argument.Board, constant.BoardFile, constant.BoardUsage)
	a.Parse(version, gitHash, buildDate)
	o := option.New()
	o.Address = a.Address()
	o.Board = board.Load(a.GetString(argument.Board))
	o.PostgresLocator = a.GetString(argument.Postgres)
	o.LitePath = a.GetString(argument.Lite)
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
