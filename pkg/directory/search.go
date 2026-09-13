package directory

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/go-ldap/ldap/v3"
)

func (c *Client) Search(
	filter string,
	attributes []string,
) ([]*Record, error) {
	connection, e := c.connect()

	if e != nil {
		return nil, e
	}

	defer errors.PanicClose(connection)
	found, f := connection.Search(
		ldap.NewSearchRequest(
			c.base,
			ldap.ScopeWholeSubtree,
			ldap.NeverDerefAliases,
			0,
			0,
			false,
			filter,
			attributes,
			nil,
		),
	)

	if f != nil {
		return nil, unexpectedSearch(f)
	}

	var result []*Record

	for _, entry := range found.Entries {
		values := map[string][]string{}

		for _, attribute := range entry.Attributes {
			values[attribute.Name] = attribute.Values
		}

		result = append(
			result,
			&Record{DistinguishedName: entry.DN, Attributes: values},
		)
	}

	return result, nil
}
