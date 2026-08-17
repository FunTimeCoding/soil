package basic

import (
	"github.com/funtimecoding/soil/pkg/opnsense/constant"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"net/http"
)

func (c *Client) Get(
	path string,
	query map[string]string,
) ([]byte, error) {
	l := locator.New(c.host).Base(constant.Base).Path(path)

	for k, v := range query {
		l.Set(k, v)
	}

	r, e := http.NewRequest(http.MethodGet, l.String(), nil)

	if e != nil {
		return nil, e
	}

	return c.send(r)
}
