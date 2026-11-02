package keepass

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/tobischo/gokeepasslib/v3"
	"os"
)

func New(
	database string,
	password string,
) (*Client, error) {
	f, e := os.Open(database)

	if e != nil {
		return nil, e
	}

	defer errors.LogClose(f)
	i, e := f.Stat()

	if e != nil {
		return nil, e
	}

	d := gokeepasslib.NewDatabase()
	d.Credentials = gokeepasslib.NewPasswordCredentials(password)

	if e = gokeepasslib.NewDecoder(f).Decode(d); e != nil {
		return nil, e
	}

	if e = d.UnlockProtectedEntries(); e != nil {
		return nil, e
	}

	return &Client{database: d, path: database, loaded: i.ModTime()}, nil
}
