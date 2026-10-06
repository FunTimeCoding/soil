package nextcloud

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"os"
)

func (c *Client) DownloadFile(
	path string,
	destination string,
) error {
	b, e := c.author.ReadFile(path)

	if e != nil {
		return e
	}

	f, g := os.Create(destination)

	if g != nil {
		return g
	}

	if _, h := f.Write(b); h != nil {
		errors.LogClose(f)

		return h
	}

	return f.Close()
}
