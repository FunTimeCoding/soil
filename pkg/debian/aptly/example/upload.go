package example

import (
	"github.com/funtimecoding/soil/pkg/debian/aptly"
	"github.com/funtimecoding/soil/pkg/errors"
)

func Upload() {
	c := aptly.New("apt.example.org", 443, false, "admin", "passwd")
	errors.PanicOnError(c.Upload("gohw-0.0.142", "gohw_0.0.142-1_amd64.deb"))
	errors.PanicOnError(c.AddToRepository("rolling", "gohw-0.0.142"))

	if false {
		errors.PanicOnError(
			c.PublishRepository(
				"rolling",
				"rolling",
				[]string{"amd64"},
				"/etc/aptly/passphrase",
			),
		)
	}

	errors.PanicOnError(c.UpdatePublish("rolling", "/etc/aptly/passphrase"))
}
