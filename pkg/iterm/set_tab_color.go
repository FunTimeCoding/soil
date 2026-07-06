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

func (c *Client) SetTabColor(
	sessionIdentifier string,
	red int,
	green int,
	blue int,
) error {
	body, e := json.Marshal(
		map[string]int{
			"red":   red,
			"green": green,
			"blue":  blue,
		},
	)

	if e != nil {
		return fmt.Errorf("set tab color: %w", e)
	}

	r, f := c.client.Post(
		c.base.Copy().Path("/sessions/%s/color", sessionIdentifier).String(),
		constant.Object,
		bytes.NewReader(body),
	)

	if f != nil {
		return fmt.Errorf("set tab color: %w", f)
	}

	defer errors.LogClose(r.Body)

	if r.StatusCode != http.StatusOK {
		b := system.ReadAll(r.Body)

		return fmt.Errorf("set tab color: %d: %s", r.StatusCode, b)
	}

	return nil
}
