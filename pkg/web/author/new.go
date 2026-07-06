package author

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/studio-b12/gowebdav"
)

func New(
	fileRoot string,
	user string,
	password string,
) *Client {
	result := gowebdav.NewClient(fileRoot, user, password)
	errors.PanicOnError(result.Connect())

	return &Client{client: result}
}
