package mock_client

import "github.com/funtimecoding/soil/pkg/firefox/tab"

func (c *Client) CreateTab(l string) (*tab.Tab, error) {
	t := tab.New()
	t.Identifier = c.nextIdentifier
	t.Locator = l
	t.Status = "complete"
	c.nextIdentifier++
	c.tabs = append(c.tabs, t)

	return t, nil
}
