package directory

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/directory/constant"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/errors/unreachable"
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
	connection, f := ldap.DialURL(address, ldap.DialWithTLSConfig(transport))

	if f != nil {
		return nil, unreachable.Format(
			"%s dial %s: %s",
			constant.Subject,
			address,
			f,
		)
	}

	if e = connection.Bind(c.bind, c.password); e != nil {
		errors.LogClose(connection)

		return nil, unreachable.Format(
			"%s service bind %s: %s",
			constant.Subject,
			c.bind,
			e,
		)
	}

	return connection, nil
}
