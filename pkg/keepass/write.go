package keepass

import (
	"errors"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/tobischo/gokeepasslib/v3"
	"os"
)

func (c *Client) write() error {
	temporary := join.Empty(c.path, ".tmp")
	f, e := os.Create(temporary)

	if e != nil {
		return e
	}

	e = errors.Join(gokeepasslib.NewEncoder(f).Encode(c.database), f.Close())

	if e == nil {
		e = os.Rename(temporary, c.path)
	}

	if e != nil {
		return errors.Join(e, os.Remove(temporary))
	}

	i, e := os.Stat(c.path)

	if e != nil {
		return e
	}

	c.loaded = i.ModTime()

	return nil
}
