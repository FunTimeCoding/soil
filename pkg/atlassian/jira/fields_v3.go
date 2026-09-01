package jira

import "github.com/funtimecoding/soil/pkg/console"

func (c *Client) FieldsV3() error {
	status, body, e := c.basic.GetPath("/rest/api/3/field")

	if e != nil {
		return e
	}

	// Does not contain more fields than the V2 API
	console.Format("Basic response: %d %s", status, body)

	return nil
}
