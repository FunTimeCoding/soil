package kestra

import (
	"context"
	"github.com/funtimecoding/soil/pkg/kestra/transport"
	"github.com/funtimecoding/soil/pkg/strings/join/key_value"
	"github.com/funtimecoding/soil/pkg/web"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/kestra-io/client-sdk/go-sdk/v2/kestra_api_client"
	"net/http"
)

// Reference: https://kestra.io/docs/api-reference/open-source
func New(
	host string,
	o ...Option,
) *Client {
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
