package aptly

import (
	"encoding/json"
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	"io"
	"net/http"
)

func (c *Client) Packages(repository string) ([]string, error) {
	r, e := c.send(
		http.MethodGet,
		fmt.Sprintf("/api/repos/%s/packages", repository),
		nil,
	)

	if e != nil {
		return nil, e
	}

	defer errors.PanicClose(r.Body)
	b, e := io.ReadAll(r.Body)

	if e != nil {
		return nil, e
	}

	var result []string
	errors.PanicOnError(json.Unmarshal(b, &result))

	return result, nil
}
