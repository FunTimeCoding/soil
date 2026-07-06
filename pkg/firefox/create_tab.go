package firefox

import (
	"encoding/json"
	"fmt"
	"github.com/funtimecoding/soil/pkg/firefox/tab"
)

func (c *Client) CreateTab(l string) (*tab.Tab, error) {
	r, e := c.send("create_tab", map[string]string{"url": l})

	if e != nil {
		return tab.Stub(), fmt.Errorf("create tab: %w", e)
	}

	var result *tab.Tab

	if f := json.Unmarshal(r.Result, result); f != nil {
		return tab.Stub(), fmt.Errorf("create tab: %w", f)
	}

	return result, nil
}
