package mock_client

import "github.com/funtimecoding/soil/pkg/backrest/response"

func (c *Client) Add(s ...*response.Stanza) {
	c.stanzas = append(c.stanzas, s...)
}
