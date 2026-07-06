package sublime

import (
	"encoding/json"
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/sublime/view"
	"github.com/funtimecoding/soil/pkg/system"
	"net/http"
)

func (c *Client) View(identifier int) (*view.View, error) {
	r, e := c.client.Get(
		c.base.Copy().Path("/views/%d", identifier).String(),
	)

	if e != nil {
		return view.Stub(), fmt.Errorf("read view: %w", e)
	}

	defer errors.LogClose(r.Body)

	if r.StatusCode != http.StatusOK {
		return view.Stub(), fmt.Errorf(
			"read view: %d: %s",
			r.StatusCode,
			system.ReadAll(r.Body),
		)
	}

	var result *view.View

	if f := json.NewDecoder(r.Body).Decode(&result); f != nil {
		return view.Stub(), fmt.Errorf("read view: %w", f)
	}

	return result, nil
}
