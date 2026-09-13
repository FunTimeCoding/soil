package directory

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/go-ldap/ldap/v3"
)

func (c *Client) Delete(distinguishedName string) error {
	connection, e := c.connect()

	if e != nil {
		return e
	}

	defer errors.PanicClose(connection)

	return classify(
		connection.Del(ldap.NewDelRequest(distinguishedName, nil)),
		distinguishedName,
	)
}
