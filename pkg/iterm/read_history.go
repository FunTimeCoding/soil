package iterm

import (
	"encoding/json"
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/iterm/history"
	"github.com/funtimecoding/soil/pkg/system"
	"net/http"
)

func (c *Client) ReadHistory(
	identifier string,
	count int,
) (*history.History, error) {
	l := c.base.Copy().Path("/sessions/%s/history", identifier).
		Set("count", fmt.Sprintf("%d", count)).String()
	r, e := c.client.Get(l)

	if e != nil {
		return history.Stub(), fmt.Errorf("read history: %w", e)
	}

	defer errors.LogClose(r.Body)

	if r.StatusCode != http.StatusOK {
		b := system.ReadAll(r.Body)

		return history.Stub(), fmt.Errorf(
			"read history: %d: %s",
			r.StatusCode,
			b,
		)
	}

	var result *history.History

	if e = json.NewDecoder(r.Body).Decode(&result); e != nil {
		return history.Stub(), fmt.Errorf("read history: %w", e)
	}

	return result, nil
}
