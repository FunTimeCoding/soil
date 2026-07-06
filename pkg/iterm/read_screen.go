package iterm

import (
	"encoding/json"
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/iterm/screen"
	"github.com/funtimecoding/soil/pkg/system"
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
		b := system.ReadAll(r.Body)

		return screen.Stub(), fmt.Errorf(
			"read screen: %d: %s",
			r.StatusCode,
			b,
		)
	}

	var result *screen.Screen

	if e = json.NewDecoder(r.Body).Decode(&result); e != nil {
		return screen.Stub(), fmt.Errorf("read screen: %w", e)
	}

	return result, nil
}
