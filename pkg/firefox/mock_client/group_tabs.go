package mock_client

import "github.com/funtimecoding/soil/pkg/firefox/tab_group"

func (c *Client) GroupTabs(
	tabIdentifiers []int,
	groupIdentifier int,
	title string,
	color string,
) (int, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if e := c.failures.Caller(); e != nil {
		return 0, e
	}

	identifier := groupIdentifier

	if identifier == 0 {
		c.groupIdentifier++
		identifier = c.groupIdentifier
	}

	c.groups[identifier] = tab_group.New(title, color)

	for _, tabIdentifier := range tabIdentifiers {
		if _, t, f := c.findTab(tabIdentifier); f == nil {
			t.GroupIdentifier = identifier
		}
	}

	return identifier, nil
}
