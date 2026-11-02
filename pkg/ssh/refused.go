package ssh

import (
	"errors"
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors/connection"
	"github.com/funtimecoding/soil/pkg/errors/constant"
	"github.com/funtimecoding/soil/pkg/errors/forbidden"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/pkg/sftp"
	"io/fs"
)

func (c *Client) refused(
	operation string,
	path string,
	e error,
) error {
	switch {
	case errors.Is(e, fs.ErrNotExist):
		return not_found.Format("%s %s: %v", operation, path, e)
	case errors.Is(e, fs.ErrPermission):
		return forbidden.Format("%s %s: %v", operation, path, e)
	case errors.Is(e, sftp.ErrSSHFxConnectionLost):
		return connection.New(constant.Dropped, c.host, "", constant.Broken)
	}

	if f := connection.Classify(e); f != nil {
		if f.Host == "" {
			return connection.New(f.Kind, c.host, f.Path, f.Reason)
		}

		return f
	}

	return fmt.Errorf("%s %s: %w", operation, path, e)
}
