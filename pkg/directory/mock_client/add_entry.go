package mock_client

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/directory/types/entry"
)

func (c *Client) AddEntry(
	unique string,
	account string,
	mail string,
	password string,
	member bool,
) {
	c.entries = append(
		c.entries,
		entry.New(
			unique,
			account,
			mail,
			account,
			fmt.Sprintf("uid=%s,ou=people", account),
		),
	)
	c.passwords[account] = password
	c.members[account] = member
}
