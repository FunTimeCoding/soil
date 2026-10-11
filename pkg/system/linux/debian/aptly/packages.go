package aptly

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/web/requester/request"
)

func (c *Client) Packages(repository string) ([]string, error) {
	var result []string

	if e := c.requester.Notation(
		request.Get(fmt.Sprintf("/api/repos/%s/packages", repository)),
		&result,
	); e != nil {
		return nil, e
	}

	return result, nil
}
