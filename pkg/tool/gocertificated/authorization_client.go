package gocertificated

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/option"
	"github.com/funtimecoding/soil/pkg/web/authorization/client"
	"github.com/funtimecoding/soil/pkg/web/constant"
)

func authorizationClient(o *option.Certificate) *client.Client {
	return client.New(
		o.Issuer,
		o.ClientIdentifier,
		o.ClientSecret,
		constant.SignInPath,
		join.Empty(o.PublicLocator, constant.CallbackPath),
		client.DeriveKey(o.EncryptionSecret),
	)
}
