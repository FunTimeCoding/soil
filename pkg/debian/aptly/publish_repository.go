package aptly

import (
	"github.com/funtimecoding/soil/pkg/debian/aptly/request"
	web "github.com/funtimecoding/soil/pkg/web/requester/request"
	"net/http"
)

func (c *Client) PublishRepository(
	repositoryName string,
	distribution string,
	architectures []string,
	passphraseFile string,
) error {
	_, e := c.requester.Bytes(
		web.New(http.MethodPost, "/api/publish/:.").WithNotation(
			request.Publish{
				SourceKind:    "local",
				Sources:       []request.PublishSource{{Name: repositoryName}},
				Architectures: architectures,
				Distribution:  distribution,
				Signing:       request.NewSignOption(passphraseFile),
			},
		),
	)

	return e
}
