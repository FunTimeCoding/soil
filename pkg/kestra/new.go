package kestra

import (
	"context"
	"github.com/funtimecoding/soil/pkg/kestra/transport"
	"github.com/funtimecoding/soil/pkg/strings/join/key_value"
	"github.com/funtimecoding/soil/pkg/web"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/kestra-io/client-sdk/go-sdk/kestra_api_client"
	"net/http"
)

func New(
	host string,
	o ...Option,
) *Client {
	// https://kestra.io/docs/how-to-guides/api
	// https://kestra.io/docs/api-reference/open-source
	// Token authentication is EE only: https://kestra.io/docs/enterprise/auth/api-tokens
	c := kestra_api_client.NewConfiguration()
	c.Scheme = constant.Secure
	c.Host = host
	result := &Client{context: context.Background()}

	for _, f := range o {
		f(result)
	}

	if result.token != "" {
		c.DefaultHeader[constant.Authorization] = key_value.Space(
			constant.Bearer,
			result.token,
		)
	}

	if result.user != "" && result.password != "" {
		l := web.Client()
		l.Transport = transport.New(
			result.user,
			result.password,
			http.DefaultTransport,
		)
		c.HTTPClient = l
	}

	result.client = kestra_api_client.NewAPIClient(c)

	return result
}
