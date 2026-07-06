package firefox

import (
	"encoding/json"
	"fmt"
	"github.com/funtimecoding/soil/pkg/firefox/content"
)

func (c *Client) ReadTab(
	identifier int,
	raw bool,
) (*content.Content, error) {
	r, e := c.send(
		"read_tab",
		map[string]any{"tab_id": identifier, "raw": raw},
	)

	if e != nil {
		return content.Stub(), fmt.Errorf("read tab: %w", e)
	}

	var result *content.Content

	if f := json.Unmarshal(r.Result, &result); f != nil {
		return content.Stub(), fmt.Errorf("read tab: %w", f)
	}

	return result, nil
}
