package claude

import (
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"os"
	"path/filepath"
)

func (c *Client) Delete(sessionIdentifier string) {
	e := os.Remove(
		filepath.Join(
			c.base,
			join.Empty(sessionIdentifier, constant.NotationLogExtension),
		),
	)

	if e != nil && !os.IsNotExist(e) {
		errors.PanicOnError(e)
	}
}
