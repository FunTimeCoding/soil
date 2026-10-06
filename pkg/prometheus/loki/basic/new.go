package basic

import (
	"github.com/funtimecoding/soil/pkg/web/locator"
	"github.com/funtimecoding/soil/pkg/web/requester"
	"github.com/funtimecoding/soil/pkg/web/requester/authorizer/basic"
)

// Reference: https://grafana.com/docs/loki/latest/reference/loki-http-api
func New(
	base *locator.Locator,
	user string,
	password string,
	verbose bool,
) *Client {
	return &Client{
		requester: requester.New(base).
			WithAuthorizer(basic.New(user, password)).
			WithRefusal(refusal),
		base:    base,
		verbose: verbose,
	}
}
