package godashboardd

import (
	"github.com/funtimecoding/go-library/pkg/strings/join"
	"github.com/funtimecoding/go-library/pkg/tool/godashboardd/constant"
	"github.com/funtimecoding/go-library/pkg/tool/godashboardd/option"
	"github.com/funtimecoding/go-library/pkg/web/authorization/client"
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
