package directory

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/go-ldap/ldap/v3"
)

func (c *Client) Add(
	distinguishedName string,
	attributes map[string][]string,
) error {
	connection, e := c.connect()

	if e != nil {
		return e
	}

	defer errors.PanicClose(connection)
	request := ldap.NewAddRequest(distinguishedName, nil)

	for name, values := range attributes {
		request.Attribute(name, values)
	}

	return classify(connection.Add(request), distinguishedName)
}
