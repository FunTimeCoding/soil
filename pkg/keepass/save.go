package keepass

import (
	"errors"
	library "github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/tobischo/gokeepasslib/v3"
	"os"
)

func (c *Client) Save() error {
	if c.Changed() {
		return errors.New(
			"database changed on disk since load - refusing to save",
		)
	}

	library.PanicOnError(c.database.LockProtectedEntries())
	temporary := join.Empty(c.path, ".tmp")
	f := system.Create(temporary)
	e := gokeepasslib.NewEncoder(f).Encode(c.database)
	library.PanicClose(f)

	if e != nil {
		return e
	}

	library.PanicOnError(os.Rename(temporary, c.path))
	library.PanicOnError(c.database.UnlockProtectedEntries())
	c.loaded = system.Stat(c.path).ModTime()

	return nil
}
