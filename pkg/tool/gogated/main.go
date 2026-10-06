package gogated

import (
	"github.com/funtimecoding/soil/pkg/argument"
	argumentConstant "github.com/funtimecoding/soil/pkg/argument/constant"
	"github.com/funtimecoding/soil/pkg/instrument"
	"github.com/funtimecoding/soil/pkg/system/environment"
	"github.com/funtimecoding/soil/pkg/tool/gogated/constant"
	"github.com/funtimecoding/soil/pkg/tool/gogated/option"
	"github.com/funtimecoding/soil/pkg/web"
)

func Main() {
	s := instrument.New(constant.Identity)
	defer func() { s.Flush(recover()) }()
	a := argument.NewInstance(constant.Identity)
	a.Web()
	a.Database()
	a.Parse()
	o := option.New()
	o.Address = a.Address()
	o.PostgresLocator = a.GetString(argumentConstant.Postgres)
	o.LitePath = a.GetString(argumentConstant.Lite)
	o.Secret = environment.Required(constant.SecretEnvironment)
	o.SuperUserMail = environment.Required(constant.SuperUserMailEnvironment)
	o.SuperUserPassword = environment.Required(
		constant.SuperUserPasswordEnvironment,
	)
	o.Issuer = environment.Required(constant.IssuerEnvironment)
	o.AdminClientIdentifier = environment.Required(
		constant.AdminClientIdentifierEnvironment,
	)
	o.AdminClientSecret = environment.Required(
		constant.AdminClientSecretEnvironment,
	)
	o.ServiceTokens = web.ServiceTokens()
	Run(o, s)
}
