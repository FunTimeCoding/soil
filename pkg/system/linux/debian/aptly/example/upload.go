package example

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/system/linux/debian/aptly"
)

func Upload() {
	c := aptly.New("apt.example.org", 443, false, "admin", "passwd")
	defer c.Close()
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
