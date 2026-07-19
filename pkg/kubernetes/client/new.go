package client

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/kubernetes/client/client_configuration"
	"github.com/funtimecoding/soil/pkg/kubernetes/client/client_context"
)

func New(clusters []string) *Client {
	result, e := FromConfiguration(
		client_configuration.New(),
		client_context.Current(),
	)
	errors.PanicOnError(e)

	if len(clusters) == 0 {
		clusters = result.ConfigurationContexts(true)
	}

	for _, l := range clusters {
		c, f := NewContext(l)

		if f != nil {
			continue
		}

		result.clients[l] = c
	}

	return result
}
