package keepass

import (
	"errors"
	"fmt"
	library "github.com/funtimecoding/soil/pkg/errors"
	"io"
	"os"
	"time"
)

func (c *Client) Backup(at time.Time) (string, error) {
	target := fmt.Sprintf("%s.%s.bak", c.path, at.Format("2006-01-02-150405"))
	source, e := os.Open(c.path)

	if e != nil {
		return "", e
	}

	defer library.LogClose(source)
	destination, e := os.Create(target)

	if e != nil {
		return "", e
	}

	_, e = io.Copy(destination, source)

	return target, errors.Join(e, destination.Close())
}
