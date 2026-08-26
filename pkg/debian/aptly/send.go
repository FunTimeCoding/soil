package aptly

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"io"
	"net/http"
)

func (c *Client) send(
	method string,
	path string,
	b io.Reader,
) (*http.Response, error) {
	r, e := http.NewRequest(method, fmt.Sprintf("%s%s", c.Base, path), b)
	errors.PanicOnError(e)

	if c.Username != "" {
		r.SetBasicAuth(c.Username, c.Password)
	}

	if b != nil {
		r.Header.Set(constant.ContentType, constant.Object)
	}

	return c.client.Do(r)
}
