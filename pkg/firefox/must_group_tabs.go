package firefox

import "github.com/funtimecoding/soil/pkg/errors"

func (c *Client) MustGroupTabs(
	tabIdentifiers []int,
	groupIdentifier int,
	title string,
	color string,
) int {
	result, e := c.GroupTabs(tabIdentifiers, groupIdentifier, title, color)
	errors.PanicOnError(e)

	return result
}
