package basic

import (
	"bytes"
	"encoding/json"
	"github.com/funtimecoding/soil/pkg/opnsense/constant"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"net/http"
)

func (c *Client) Post(
	path string,
	body any,
) ([]byte, error) {
	j, e := json.Marshal(body)

	if e != nil {
		return nil, e
	}

	r, f := http.NewRequest(
		http.MethodPost,
		locator.New(c.host).Base(constant.Base).Path(path).String(),
		bytes.NewReader(j),
	)

	if f != nil {
		return nil, f
	}

	r.Header.Set(webConstant.ContentType, webConstant.Object)

	return c.send(r)
}
