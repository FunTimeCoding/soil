package author

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/web"
	"github.com/studio-b12/gowebdav"
)

func New(
	fileRoot string,
	user string,
	password string,
) *Client {
	result := gowebdav.NewClient(fileRoot, user, password)
	result.SetTransport(web.StallClient().Transport)
	errors.PanicOnError(result.Connect())

	return &Client{client: result}
}
