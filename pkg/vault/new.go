package vault

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"github.com/hashicorp/vault-client-go"
	"time"
)

func New(host string) *Client {
	client, e := vault.New(
		vault.WithAddress(locator.New(host).String()),
		vault.WithRequestTimeout(10*time.Second),
	)
	errors.PanicOnError(e)

	return &Client{client: client}
}
