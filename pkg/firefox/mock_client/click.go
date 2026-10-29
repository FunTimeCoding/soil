package mock_client

import "github.com/funtimecoding/soil/pkg/firefox/element"

func (c *Client) Click(
	identifier int,
	selector string,
) (*element.Element, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if e := c.failures.Caller(); e != nil {
		return element.Stub(), e
	}

	if _, _, f := c.findTab(identifier); f != nil {
		return element.Stub(), f
	}

	result := element.New()
	result.Selector = selector
	result.Tag = "button"

	return result, nil
}
