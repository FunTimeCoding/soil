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

func (c *Client) SendText(
	identifier string,
	text string,
) error {
	body, e := json.Marshal(map[string]string{"text": text})

	if e != nil {
		return fmt.Errorf("send text: %w", e)
	}

	r, f := c.client.Post(
		c.base.Copy().Path("/sessions/%s/send", identifier).String(),
		constant.Object,
		bytes.NewReader(body),
	)

	if f != nil {
		return fmt.Errorf("send text: %w", f)
	}

	defer errors.LogClose(r.Body)

	if r.StatusCode != http.StatusOK {
		return fmt.Errorf(
			"send text: %d: %s",
			r.StatusCode,
			system.ReadAll(r.Body),
		)
	}

	return nil
}
