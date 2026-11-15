package counted_directory

import "github.com/funtimecoding/soil/pkg/directory/types/entry"

func (c *Directory) Authenticate(
	account string,
	password string,
) (*entry.Entry, error) {
	c.Authentications++

	return c.Client.Authenticate(account, password)
}
