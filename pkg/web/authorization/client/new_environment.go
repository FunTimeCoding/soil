package client

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/system/environment"
	"github.com/funtimecoding/soil/pkg/web/constant"
)

func NewEnvironment() *Client {
	return New(
		environment.Required(constant.AuthorizationIssuerEnvironment),
		environment.Required(constant.AuthorizationClientIdentifierEnvironment),
		environment.Required(constant.AuthorizationClientSecretEnvironment),
		constant.SignInPath,
		join.Empty(
			environment.Required(constant.PublicLocatorEnvironment),
			constant.CallbackPath,
		),
		DeriveKey(
			environment.Required(
				constant.AuthorizationEncryptionSecretEnvironment,
			),
		),
	)
}
