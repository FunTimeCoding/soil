package mock_client

import "github.com/funtimecoding/soil/pkg/postgres/types/statement"

func (c *Client) Statements() []*statement.Statement {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	return c.statements
}
