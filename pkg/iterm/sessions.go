package iterm

import (
	"encoding/json"
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/iterm/response"
	"github.com/funtimecoding/soil/pkg/iterm/session"
	"github.com/funtimecoding/soil/pkg/system"
	"net/http"
)

func (c *Client) Sessions() ([]*session.Session, error) {
	l := c.base.Copy().Path("/sessions").String()
	r, e := c.client.Get(l)

	if e != nil {
		return nil, fmt.Errorf("list sessions: %w", e)
	}

	defer errors.LogClose(r.Body)

	if r.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(
			"list sessions: %d: %s",
			r.StatusCode,
			system.ReadAll(r.Body),
		)
	}

	var result *response.Sessions

	if e = json.NewDecoder(r.Body).Decode(&result); e != nil {
		return nil, fmt.Errorf("list sessions: %w", e)
	}

	return result.Sessions, nil
}
