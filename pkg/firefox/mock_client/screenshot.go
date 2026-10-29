package mock_client

import (
	"encoding/base64"
	"github.com/funtimecoding/soil/pkg/firefox/capture"
)

func (c *Client) Screenshot(identifier int) (*capture.Capture, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if e := c.failures.Caller(); e != nil {
		return capture.Stub(), e
	}

	_, t, f := c.findTab(identifier)

	if f != nil {
		return capture.Stub(), f
	}

	result := capture.New()
	result.Identifier = identifier
	result.Locator = t.Locator
	result.Title = t.Title
	result.Content = base64.StdEncoding.EncodeToString([]byte("mock image"))

	return result, nil
}
