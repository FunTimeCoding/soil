package iterm

import (
	"encoding/json"
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/iterm/session"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"net/http"
)

func (c *Client) CreateTab() (*session.Session, error) {
	r, e := c.client.Post(
		c.base.Copy().Path("/tabs").String(),
		constant.Object,
		nil,
	)

	if e != nil {
		return session.Stub(), fmt.Errorf("create tab: %w", e)
	}

	defer errors.LogClose(r.Body)

	if r.StatusCode != http.StatusCreated {
		return session.Stub(), fmt.Errorf(
			"create tab: %d: %s",
			r.StatusCode,
			system.ReadAll(r.Body),
		)
	}

	var result *session.Session

	if f := json.NewDecoder(r.Body).Decode(&result); f != nil {
		return session.Stub(), fmt.Errorf("create tab: %w", f)
	}

	return result, nil
}
