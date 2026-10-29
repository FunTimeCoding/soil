package mock_client

import "github.com/funtimecoding/soil/pkg/firefox/element"

func (c *Client) Fill(
	identifier int,
	selector string,
	value string,
	_ bool,
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
	result.Tag = "input"
	result.Text = value

	return result, nil
}
