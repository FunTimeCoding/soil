package iterm

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"net/http"
)

func (c *Client) SetTabTitle(
	tabIdentifier string,
	title string,
) error {
	body, e := json.Marshal(map[string]string{"title": title})

	if e != nil {
		return fmt.Errorf("set tab title: %w", e)
	}

	r, f := c.client.Post(
		c.base.Copy().Path("/tabs/%s/title", tabIdentifier).String(),
		constant.Object,
		bytes.NewReader(body),
	)

	if f != nil {
		return fmt.Errorf("set tab title: %w", f)
	}

	defer errors.LogClose(r.Body)

	if r.StatusCode != http.StatusOK {
		return fmt.Errorf(
			"set tab title: %d: %s",
			r.StatusCode,
			system.ReadAll(r.Body),
		)
	}

	return nil
}
