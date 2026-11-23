package graph

import (
	"github.com/funtimecoding/soil/pkg/microsoft/constant"
	"github.com/funtimecoding/soil/pkg/microsoft/keychain_store"
	"github.com/funtimecoding/soil/pkg/system/environment"
	"github.com/funtimecoding/soil/pkg/web/locator"
)

func NewEnvironment() *Client {
	client := environment.Required(constant.ClientEnvironment)

	return New(
		locator.New(constant.GraphHost),
		locator.New(constant.LoginHost),
		environment.Required(constant.TenantEnvironment),
		client,
		keychain_store.New(constant.KeychainService, client),
	)
}
