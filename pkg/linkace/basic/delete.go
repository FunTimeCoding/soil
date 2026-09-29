package basic

import (
	"encoding/json"
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/web"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/detail_error"
	"net/http"
)

func (c *Client) Delete(path string) error {
	l := c.base.Copy().Path(path)
	r := web.NewDelete(l.String())
	web.Bearer(r, c.token)
	r.Header.Set(constant.Accept, constant.Object)
	r.Header.Set(constant.ContentType, constant.Object)
	response := web.Send(web.Client(), r)
	defer errors.PanicClose(response.Body)

	if response.StatusCode != http.StatusOK &&
		response.StatusCode != http.StatusNoContent {
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

	return nil
}
