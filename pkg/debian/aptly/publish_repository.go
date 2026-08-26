package aptly

import (
	"bytes"
	"github.com/funtimecoding/soil/pkg/debian/aptly/request"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/errors/unexpected"
	"github.com/funtimecoding/soil/pkg/notation"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"net/http"
)

func (c *Client) PublishRepository(
	repositoryName string,
	distribution string,
	architectures []string,
	passphraseFile string,
) error {
	r, e := c.send(
		constant.Post,
		"/api/publish/:.",
		bytes.NewReader(
			notation.Marshal(
				request.Publish{
					SourceKind: "local",
					Sources: []request.PublishSource{
						{Name: repositoryName},
					},
					Architectures: architectures,
					Distribution:  distribution,
					Signing:       request.NewSignOption(passphraseFile),
				},
			),
		),
	)
	errors.PanicOnError(e)
	defer errors.LogClose(r.Body)

	if r.StatusCode != http.StatusCreated &&
		r.StatusCode != http.StatusOK {
		return unexpected.Format(
			"publish status: %s %s",
			r.Status,
			string(system.ReadAll(r.Body)),
		)
	}

	return nil
}
