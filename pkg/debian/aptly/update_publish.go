package aptly

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/debian/aptly/request"
	web "github.com/funtimecoding/soil/pkg/web/requester/request"
	"net/http"
)

func (c *Client) UpdatePublish(
	distribution string,
	passphraseFile string,
) error {
	_, e := c.requester.Bytes(
		web.New(
			http.MethodPut,
			fmt.Sprintf("/api/publish/:./%s", distribution),
		).WithNotation(
			request.UpdatePublish{
				Signing: request.NewSignOption(passphraseFile),
			},
		),
	)

	return e
}
