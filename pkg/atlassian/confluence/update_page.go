package confluence

import "github.com/funtimecoding/soil/pkg/atlassian/confluence/page"

func (c *Client) UpdatePage(
	identifier string,
	title string,
	markdown string,
	message string,
) (*page.Page, error) {
	current, e := c.Page(identifier)

	if e != nil {
		return nil, e
	}

	return c.UpdatePageAt(
		identifier,
		title,
		markdown,
		current.Raw.Version.Number+1,
		message,
	)
}
