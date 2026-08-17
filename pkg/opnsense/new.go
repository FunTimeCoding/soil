package opnsense

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/opnsense/basic"
)

func New(
	host string,
	key string,
	secret string,
	insecure bool,
) *Client {
	errors.FatalOnEmpty(host, "host")
	errors.FatalOnEmpty(key, "key")
	errors.FatalOnEmpty(secret, "secret")

	return &Client{basic: basic.New(host, key, secret, insecure)}
}
