package keepass

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/tobischo/gokeepasslib/v3"
)

func New(
	database string,
	password string,
) *Client {
	f := system.Open(database)
	d := gokeepasslib.NewDatabase()
	d.Credentials = gokeepasslib.NewPasswordCredentials(password)
	errors.PanicOnError(gokeepasslib.NewDecoder(f).Decode(d))
	errors.PanicOnError(d.UnlockProtectedEntries())

	return &Client{database: d}
}
