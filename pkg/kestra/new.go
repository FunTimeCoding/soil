package kestra

import (
	"context"
	"github.com/funtimecoding/soil/pkg/strings/join/key_value"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/stall_client"
	"github.com/kestra-io/client-sdk/go-sdk/v2/kestra_api_client"
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

	b := stall_client.New()

	if result.user != "" && result.password != "" {
		b.BasicAuth(result.user, result.password)
	}

	result.pool = b.Build()
	c.HTTPClient = result.pool.Client()
	result.client = kestra_api_client.NewAPIClient(c)

	return result
}
