package firefox

import (
	"encoding/json"
	"fmt"
	"github.com/funtimecoding/soil/pkg/firefox/capture"
)

func (c *Client) Screenshot(identifier int) (*capture.Capture, error) {
	r, e := c.send("screenshot", map[string]any{"tab_id": identifier})

	if e != nil {
		return capture.Stub(), fmt.Errorf("screenshot: %w", e)
	}

	var result *capture.Capture

	if f := json.Unmarshal(r.Result, &result); f != nil {
		return capture.Stub(), fmt.Errorf("screenshot: %w", f)
	}

	return result, nil
}
