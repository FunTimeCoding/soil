package directory

import "github.com/funtimecoding/soil/pkg/errors"

func (c *Client) InGroup(account string) (bool, error) {
	connection, e := c.connect()

	if e != nil {
		return false, e
	}

	defer errors.PanicClose(connection)
	result, e := connection.Search(c.searchRequest(c.groupFilter, account))

	if e != nil {
		return false, unexpectedSearch(e)
	}

	return len(result.Entries) > 0, nil
}
