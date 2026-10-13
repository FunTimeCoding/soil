package mock_client

import (
	"encoding/base64"
	"fmt"
	"github.com/funtimecoding/soil/pkg/firefox/capture"
)

func (c *Client) Screenshot(identifier int) (*capture.Capture, error) {
	for _, t := range c.tabs {
		if t.Identifier == identifier {
			result := capture.New()
			result.Identifier = identifier
			result.Locator = t.Locator
			result.Title = t.Title
			result.Content = base64.StdEncoding.EncodeToString(
				[]byte("mock image"),
			)

			return result, nil
		}
	}

	return capture.Stub(), fmt.Errorf("tab %d not found", identifier)
}
