package directory

import (
	"github.com/funtimecoding/soil/pkg/directory/types/search_record"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/go-ldap/ldap/v3"
)

func (c *Client) Search(
	filter string,
	attributes []string,
) ([]*search_record.Record, error) {
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

	var result []*search_record.Record

	for _, entry := range found.Entries {
		values := map[string][]string{}

		for _, attribute := range entry.Attributes {
			values[attribute.Name] = attribute.Values
		}

		result = append(result, search_record.New(entry.DN, values))
	}

	return result, nil
}
