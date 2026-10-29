package directory

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/directory/constant"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/errors/connection"
	"github.com/go-ldap/ldap/v3"
)

func (c *Client) connect() (*ldap.Conn, error) {
	scheme := "ldaps"

	if c.insecure {
		scheme = "ldap"
	}

	transport, e := c.transport()

	if e != nil {
		return nil, e
	}

	address := fmt.Sprintf("%s://%s:%d", scheme, c.host, c.port)
	link, f := ldap.DialURL(address, ldap.DialWithTLSConfig(transport))

	if f != nil {
		if g := connection.Classify(f); g != nil {
			return nil, g
		}

		return nil, fmt.Errorf("%s dial %s: %w", constant.Subject, address, f)
	}

	if g := link.Bind(c.bind, c.password); g != nil {
		errors.LogClose(link)

		return nil, classify(g, c.bind)
	}

	return link, nil
}
