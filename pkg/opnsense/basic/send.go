package basic

import (
	"github.com/funtimecoding/soil/pkg/web"
	"io"
	"net/http"
)

func (c *Client) send(r *http.Request) ([]byte, error) {
	r.SetBasicAuth(c.key, c.secret)
	client := web.Client()

	if c.insecure {
		client = web.InsecureClient()
	}

	s, e := client.Do(r)

	if e != nil {
		return nil, e
	}

	result, f := io.ReadAll(s.Body)
	g := s.Body.Close()

	if f != nil {
		return nil, f
	}

	if g != nil {
		return nil, g
	}

	if s.StatusCode >= http.StatusBadRequest {
		return nil, parseDetail(result, s.Status)
	}

	return result, nil
}
