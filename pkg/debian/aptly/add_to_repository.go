package aptly

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/errors/unexpected"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"net/http"
)

func (c *Client) AddToRepository(
	repoName string,
	directory string,
) error {
	r, e := c.send(
		constant.Post,
		fmt.Sprintf("/api/repos/%s/file/%s", repoName, directory),
		nil,
	)
	errors.PanicOnError(e)
	defer errors.LogClose(r.Body)

	if r.StatusCode != http.StatusOK {
		return unexpected.Format(
			"add status: %s %s",
			r.Status,
			string(system.ReadAll(r.Body)),
		)
	}

	return nil
}
