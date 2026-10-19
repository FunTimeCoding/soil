package missing_link

import "github.com/funtimecoding/soil/pkg/errors/not_found"

func (c *Client) DeleteLink(identifier int) error {
	return not_found.Format("resource not found: links/%d", identifier)
}
