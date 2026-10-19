package iterm

import (
	"encoding/json"
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/iterm/screen"
	"net/http"
)

func (c *Client) ReadScreen(identifier string) (*screen.Screen, error) {
	r, e := c.client.Get(
		c.base.Copy().Path("/sessions/%s/screen", identifier).String(),
	)

	if e != nil {
		return screen.Stub(), fmt.Errorf("read screen: %w", e)
	}

	defer errors.LogClose(r.Body)

	if r.StatusCode != http.StatusOK {
		return screen.Stub(), statusError("read screen", r)
	}

	var result *screen.Screen

	if e = json.NewDecoder(r.Body).Decode(&result); e != nil {
		return screen.Stub(), fmt.Errorf("read screen: %w", e)
	}

	return result, nil
}
