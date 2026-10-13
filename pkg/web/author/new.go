package author

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/web/stall_client"
	"github.com/studio-b12/gowebdav"
)

func New(
	fileRoot string,
	user string,
	password string,
) *Client {
	p := stall_client.New().Build()
	result := gowebdav.NewClient(fileRoot, user, password)
	result.SetTransport(p.Client().Transport)
	errors.PanicOnError(result.Connect())

	return &Client{client: result, pool: p}
}
