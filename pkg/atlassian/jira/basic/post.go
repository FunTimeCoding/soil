package basic

import (
	"github.com/funtimecoding/soil/pkg/web/constant"
	"io"
	"net/http"
	"strings"
)

func (c *Client) Post(
	l string,
	body string,
) (int, string, error) {
	r, e := http.NewRequest(
		http.MethodPost,
		l,
		strings.NewReader(body),
	)

	if e != nil {
		return 0, "", e
	}

	r.SetBasicAuth(c.user, c.token)
	r.Header.Add(constant.ContentType, constant.Object)
	r.Header.Add(constant.Accept, constant.Object)
	result, f := http.DefaultClient.Do(r)

	if f != nil {
		return 0, "", f
	}

	b, g := io.ReadAll(result.Body)

	if h := result.Body.Close(); h != nil {
		return 0, "", h
	}

	if g != nil {
		return 0, "", g
	}

	return result.StatusCode, string(b), nil
}
