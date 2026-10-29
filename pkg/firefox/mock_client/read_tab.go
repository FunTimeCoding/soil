package mock_client

import "github.com/funtimecoding/soil/pkg/firefox/content"

func (c *Client) ReadTab(
	identifier int,
	_ bool,
) (*content.Content, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if e := c.failures.Caller(); e != nil {
		return content.Stub(), e
	}

	_, t, f := c.findTab(identifier)

	if f != nil {
		return content.Stub(), f
	}

	result := content.New()
	result.Identifier = t.Identifier
	result.Locator = t.Locator
	result.Title = t.Title

	return result, nil
}
