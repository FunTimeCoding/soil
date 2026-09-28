package jellyfin

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func (c *Client) get(
	path string,
	v url.Values,
	out any,
) error {
	u := fmt.Sprintf("%s%s", c.base, path)

	if len(v) > 0 {
		u = fmt.Sprintf("%s?%s", u, v.Encode())
	}

	r, e := c.http.Get(u)

	if e != nil {
		return e
	}

	b, f := io.ReadAll(r.Body)
	g := r.Body.Close()

	if f != nil {
		return f
	}

	if g != nil {
		return g
	}

	if r.StatusCode != http.StatusOK {
		return parseDetail(b, r.Status)
	}

	return json.Unmarshal(b, out)
}
