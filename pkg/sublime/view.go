package sublime

import (
	"encoding/json"
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/sublime/view"
	"net/http"
)

func (c *Client) View(identifier int) (*view.View, error) {
	r, e := c.client.Get(c.base.Copy().Path("/views/%d", identifier).String())

	if e != nil {
		return view.Stub(), fmt.Errorf("read view: %w", e)
	}

	defer errors.LogClose(r.Body)

	if r.StatusCode != http.StatusOK {
		return view.Stub(), statusError("read view", r)
	}

	var result *view.View

	if f := json.NewDecoder(r.Body).Decode(&result); f != nil {
		return view.Stub(), fmt.Errorf("read view: %w", f)
	}

	return result, nil
}
