package directory

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/go-ldap/ldap/v3"
)

func (c *Client) Modify(
	distinguishedName string,
	attributes map[string][]string,
) error {
	connection, e := c.connect()

	if e != nil {
		return e
	}

	defer errors.PanicClose(connection)
	request := ldap.NewModifyRequest(distinguishedName, nil)

	for name, values := range attributes {
		if len(values) == 0 {
			request.Delete(name, nil)

			continue
		}

		request.Replace(name, values)
	}

	return classify(connection.Modify(request), distinguishedName)
}
