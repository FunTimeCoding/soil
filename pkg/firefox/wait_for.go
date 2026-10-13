package firefox

import (
	"encoding/json"
	"fmt"
	"github.com/funtimecoding/soil/pkg/firefox/tab"
)

func (c *Client) WaitFor(
	identifier int,
	selector string,
	timeout int,
) (*tab.Tab, error) {
	parameters := map[string]any{"tab_id": identifier}

	if selector != "" {
		parameters["selector"] = selector
	}

	if timeout > 0 {
		parameters["timeout"] = timeout
	}

	r, e := c.send("wait_for", parameters)

	if e != nil {
		return tab.Stub(), fmt.Errorf("wait for: %w", e)
	}

	var result *tab.Tab

	if f := json.Unmarshal(r.Result, &result); f != nil {
		return tab.Stub(), fmt.Errorf("wait for: %w", f)
	}

	return result, nil
}
