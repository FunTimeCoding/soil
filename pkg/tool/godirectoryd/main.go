package godirectoryd

import (
	"github.com/funtimecoding/soil/pkg/argument"
	"github.com/funtimecoding/soil/pkg/instrument"
	"github.com/funtimecoding/soil/pkg/system/environment"
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/option"
	"github.com/funtimecoding/soil/pkg/web"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
)

func Main(
	version string,
	gitHash string,
	buildDate string,
) {
	i := instrument.New(constant.Identity, version)
	defer func() { i.Flush(recover()) }()
	a := argument.NewInstance(constant.Identity)
	a.Web()
	a.Parse(version, gitHash, buildDate)
	o := option.New()
	o.Address = a.Address()
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
	o.Version = version
	Run(o, i)
}
