package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/kubernetes/client/client"
	"github.com/funtimecoding/soil/pkg/kubernetes/client/client_configuration"
	"github.com/funtimecoding/soil/pkg/kubernetes/client/client_context"
	"github.com/funtimecoding/soil/pkg/kubernetes/client/metrics"
)

func New(clusters []string) *Client {
	result := Stub()
	result.context = context.Background()
	result.configuration = client_configuration.New()
	result.client = client.New(result.configuration)
	result.metric = metrics.New(result.configuration)
	result.cluster = client_context.Current()

	if len(clusters) == 0 {
		clusters = result.ConfigurationContexts(true)
	}

	for _, l := range clusters {
		c, e := NewContext(l)

		if e != nil {
			continue
		}

		result.clients[l] = c
	}

	return result
}
