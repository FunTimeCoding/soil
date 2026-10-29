package netbox

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/netbox/cache"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"github.com/funtimecoding/soil/pkg/web/stall_client"
	"github.com/netbox-community/go-netbox/v4"
)

func New(
	l *locator.Locator,
	token string,
	o ...Option,
) *Client {
	errors.FatalOnEmpty(token, "token")
	result := &Client{
		context: context.Background(),
		client:  netbox.NewAPIClientFor(l.String(), token),
		base:    l,
		cache:   cache.New(),
	}

	for _, f := range o {
		f(result)
	}

	b := stall_client.New()

	if result.observer != nil {
		b.Observe(result.observer)
	}

	result.pool = b.Build()
	result.client.GetConfig().HTTPClient = result.pool.Client()

	if result.verbose {
		result.client.GetConfig().Debug = true
	}

	return result
}
