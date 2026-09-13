package directory

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/go-ldap/ldap/v3"
)

func (c *Client) SetPassword(
	distinguishedName string,
	password string,
) error {
	connection, e := c.connect()

	if e != nil {
		return e
	}

	defer errors.PanicClose(connection)
	_, f := connection.PasswordModify(
		ldap.NewPasswordModifyRequest(distinguishedName, "", password),
	)

	return classify(f, distinguishedName)
}
