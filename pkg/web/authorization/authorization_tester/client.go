package authorization_tester

import (
	"github.com/funtimecoding/soil/pkg/web/authorization/client"
	"github.com/funtimecoding/soil/pkg/web/constant"
)

func Client() *client.Client {
	return client.New(
		"https://issuer.test",
		"",
		"",
		constant.SignInPath,
		"https://service.test/callback",
		client.DeriveKey("secret"),
	)
}
