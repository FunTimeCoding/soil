package basic

import (
	atlassian "github.com/funtimecoding/soil/pkg/atlassian/constant"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"io"
	"net/http"
	"strings"
)

func (c *Client) PostOldPath(
	p string,
	body string,
) (string, error) {
	r, e := http.NewRequest(
		http.MethodPost,
		locator.New(c.host).Base(atlassian.ConfluenceOldBase).Path(p).String(),
		strings.NewReader(body),
	)

	if e != nil {
		return "", e
	}

	r.SetBasicAuth(c.user, c.token)
	r.Header[webConstant.Accept] = []string{webConstant.Object}
	r.Header[webConstant.ContentType] = []string{webConstant.Object}
	result, f := http.DefaultClient.Do(r)

	if f != nil {
		return "", f
	}

	b, g := io.ReadAll(result.Body)

	if h := result.Body.Close(); h != nil {
		return "", h
	}

	if g != nil {
		return "", g
	}

	if result.StatusCode >= http.StatusBadRequest {
		return "", parseDetail(b, result.Status)
	}

	return string(b), nil
}
