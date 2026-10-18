package client

import "github.com/funtimecoding/soil/pkg/tool/gomemoryd/generated/client"

func NewRestClient(
	c *client.ClientWithResponses,
	address string,
) *RestClient {
	return &RestClient{http: c, address: address}
}
