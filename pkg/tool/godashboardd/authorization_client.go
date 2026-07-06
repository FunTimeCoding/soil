package godashboardd

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/godashboardd/constant"
	"github.com/funtimecoding/soil/pkg/tool/godashboardd/option"
	"github.com/funtimecoding/soil/pkg/web/authorization/client"
)

func authorizationClient(o *option.Dashboard) *client.Client {
	if o.Issuer == "" {
		return nil
	}

	return client.New(
		o.Issuer,
		o.ClientIdentifier,
		o.ClientSecret,
		constant.SignInPath,
		join.Empty(o.PublicLocator, constant.CallbackPath),
		client.DeriveKey(o.EncryptionSecret),
	)
}
