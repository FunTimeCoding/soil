package basic

import (
	"encoding/json"
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/web"
	"github.com/funtimecoding/soil/pkg/web/detail_error"
	"net/http"
)

func (c *Client) Get(
	path string,
	params map[string]string,
	out any,
) error {
	l := c.base.Copy().Path(path)

	for k, v := range params {
		l.Set(k, v)
	}

	r := web.NewGet(l.String())
	web.Bearer(r, c.token)
	r.Header.Set("Accept", "application/json")
	response := web.Send(web.Client(), r)
	defer errors.PanicClose(response.Body)

	if response.StatusCode != http.StatusOK {
		var check errorResponse

		if json.NewDecoder(response.Body).Decode(&check) == nil &&
			check.Message != "" {
			return detail_error.New(check.Message, response.Status)
		}

		return detail_error.New(
			fmt.Sprintf("%s %s", response.Status, l.String()),
			response.Status,
		)
	}

	return json.NewDecoder(response.Body).Decode(out)
}
