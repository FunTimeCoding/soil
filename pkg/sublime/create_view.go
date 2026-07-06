package sublime

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/sublime/view"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"net/http"
)

func (c *Client) CreateView(
	title string,
	content string,
	syntax string,
) (*view.View, error) {
	body, e := json.Marshal(
		map[string]string{
			"title":   title,
			"content": content,
			"syntax":  syntax,
		},
	)

	if e != nil {
		return view.Stub(), fmt.Errorf("create view: %w", e)
	}

	r, f := c.client.Post(
		c.base.Copy().Path("/views").String(),
		constant.Object,
		bytes.NewReader(body),
	)

	if f != nil {
		return view.Stub(), fmt.Errorf("create view: %w", f)
	}

	defer errors.LogClose(r.Body)

	if r.StatusCode != http.StatusCreated {
		return view.Stub(), fmt.Errorf(
			"create view: %d: %s",
			r.StatusCode,
			system.ReadAll(r.Body),
		)
	}

	var result *view.View

	if g := json.NewDecoder(r.Body).Decode(&result); g != nil {
		return view.Stub(), fmt.Errorf("create view: %w", g)
	}

	return result, nil
}
