package firefox

import (
	"encoding/json"
	"fmt"
	"github.com/funtimecoding/soil/pkg/firefox/tab"
)

func (c *Client) Tabs() ([]*tab.Tab, error) {
	r, e := c.send("list_tabs", nil)

	if e != nil {
		return nil, fmt.Errorf("list tabs: %w", e)
	}

	var result []*tab.Tab

	if f := json.Unmarshal(r.Result, &result); f != nil {
		return nil, fmt.Errorf("list tabs: %w", f)
	}

	return result, nil
}
