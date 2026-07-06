package client

import "github.com/funtimecoding/soil/pkg/tool/gomemoryd/generated/client"

func NewRestClient(c *client.Client) *RestClient {
	return &RestClient{http: c}
}
