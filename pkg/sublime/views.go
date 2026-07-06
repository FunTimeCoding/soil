package sublime

import (
	"encoding/json"
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/sublime/response"
	"github.com/funtimecoding/soil/pkg/sublime/view"
	"github.com/funtimecoding/soil/pkg/system"
	"net/http"
)

func (c *Client) Views() ([]*view.View, error) {
	r, e := c.client.Get(c.base.Copy().Path("/views").String())

	if e != nil {
		return nil, fmt.Errorf("list views: %w", e)
	}

	defer errors.LogClose(r.Body)

	if r.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(
			"list views: %d: %s",
			r.StatusCode,
			system.ReadAll(r.Body),
		)
	}

	var result *response.Views

	if e = json.NewDecoder(r.Body).Decode(&result); e != nil {
		return nil, fmt.Errorf("list views: %w", e)
	}

	return result.Views, nil
}
