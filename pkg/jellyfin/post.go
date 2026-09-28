package jellyfin

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"io"
	"net/http"
	"net/url"
)

func (c *Client) post(
	path string,
	v url.Values,
	body any,
) error {
	u := fmt.Sprintf("%s%s", c.base, path)

	if len(v) > 0 {
		u = fmt.Sprintf("%s?%s", u, v.Encode())
	}

	var payload []byte

	if body != nil {
		var e error
		payload, e = json.Marshal(body)

		if e != nil {
			return e
		}
	}

	request, e := http.NewRequest(http.MethodPost, u, bytes.NewReader(payload))

	if e != nil {
		return e
	}

	if body != nil {
		request.Header.Set(constant.ContentType, constant.Object)
	}

	r, f := c.http.Do(request)

	if f != nil {
		return f
	}

	b, g := io.ReadAll(r.Body)
	h := r.Body.Close()

	if g != nil {
		return g
	}

	if h != nil {
		return h
	}

	if r.StatusCode != http.StatusNoContent &&
		r.StatusCode != http.StatusOK {
		return parseDetail(b, r.Status)
	}

	return nil
}
