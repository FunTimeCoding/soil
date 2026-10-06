package unit

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/web/authorization/client"
	"github.com/funtimecoding/soil/pkg/web/constant"
)

func newAuthorizationClient(
	issuer string,
	r *memory.Memory,
) *client.Client {
	return client.New(
		issuer,
		"delta",
		"hotel",
		constant.SignInPath,
		"http://app.test/callback",
		client.DeriveKey("foxtrot"),
	).WithReporter(r)
}
