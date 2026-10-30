package aptly

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors/conflict"
	"github.com/funtimecoding/soil/pkg/system/constant"
	"github.com/funtimecoding/soil/pkg/system/linux/debian/aptly/response"
	"github.com/funtimecoding/soil/pkg/web/requester/request"
	"net/http"
	"strings"
)

func (c *Client) AddToRepository(
	repository string,
	directory string,
	force bool,
) error {
	q := request.New(
		http.MethodPost,
		fmt.Sprintf("/api/repos/%s/file/%s", repository, directory),
	)

	if force {
		q.WithParameter(
			constant.AptlyForceReplaceParameter,
			constant.AptlyForceReplaceValue,
		)
	}

	var result response.Add

	if e := c.requester.Notation(q, &result); e != nil {
		return e
	}

	if len(result.FailedFiles) > 0 {
		return conflict.Format(
			"aptly refused %s: %s",
			strings.Join(result.FailedFiles, ", "),
			strings.Join(result.Report.Warnings, "; "),
		)
	}

	return nil
}
