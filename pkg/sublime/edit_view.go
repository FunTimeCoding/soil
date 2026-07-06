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

func (c *Client) EditView(
	identifier int,
	old string,
	new string,
	all bool,
) (*view.View, error) {
	body, e := json.Marshal(
		map[string]any{
			"old_string":  old,
			"new_string":  new,
			"replace_all": all,
		},
	)

	if e != nil {
		return view.Stub(), fmt.Errorf("edit view: %w", e)
	}

	q, f := http.NewRequest(
		http.MethodPut,
		c.base.Copy().Path("/views/%d", identifier).String(),
		bytes.NewReader(body),
	)

	if f != nil {
		return view.Stub(), fmt.Errorf("edit view: %w", f)
	}

	q.Header.Set(constant.ContentType, constant.Object)
	r, g := c.client.Do(q)

	if g != nil {
		return view.Stub(), fmt.Errorf("edit view: %w", g)
	}

	defer errors.LogClose(r.Body)

	if r.StatusCode != http.StatusOK {
		return view.Stub(), fmt.Errorf(
			"edit view: %d: %s",
			r.StatusCode,
			system.ReadAll(r.Body),
		)
	}

	var result *view.View

	if h := json.NewDecoder(r.Body).Decode(&result); h != nil {
		return view.Stub(), fmt.Errorf("edit view: %w", h)
	}

	return result, nil
}
