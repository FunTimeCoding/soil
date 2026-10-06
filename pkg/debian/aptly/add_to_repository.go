package aptly

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/web/requester/request"
	"net/http"
)

func (c *Client) AddToRepository(
	repoName string,
	directory string,
) error {
	_, e := c.requester.Bytes(
		request.New(
			http.MethodPost,
			fmt.Sprintf("/api/repos/%s/file/%s", repoName, directory),
		),
	)

	return e
}
