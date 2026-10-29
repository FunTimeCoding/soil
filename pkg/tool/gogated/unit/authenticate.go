package unit

import "github.com/funtimecoding/soil/pkg/directory/types/entry"

func (c *countedDirectory) Authenticate(
	account string,
	password string,
) (*entry.Entry, error) {
	c.authentications++

	return c.Client.Authenticate(account, password)
}
