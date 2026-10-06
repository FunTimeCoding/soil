package opnsense

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/opnsense/basic"
	"github.com/funtimecoding/soil/pkg/opnsense/constant"
	"github.com/funtimecoding/soil/pkg/web/locator"
)

func New(
	host string,
	key string,
	secret string,
	untrusted bool,
) *Client {
	errors.FatalOnEmpty(host, "host")
	errors.FatalOnEmpty(key, "key")
	errors.FatalOnEmpty(secret, "secret")

	return &Client{
		basic: basic.New(
			locator.New(host).Base(constant.Base),
			key,
			secret,
			untrusted,
		),
	}
}
