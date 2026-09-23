package mock_client

import (
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"os"
	"path/filepath"
)

func (c *Client) Delete(sessionIdentifier string) {
	if c.Harbor == "" {
		return
	}

	e := os.Remove(
		filepath.Join(
			c.Harbor,
			join.Empty(sessionIdentifier, constant.NotationLogExtension),
		),
	)

	if e != nil && !os.IsNotExist(e) {
		errors.PanicOnError(e)
	}
}
