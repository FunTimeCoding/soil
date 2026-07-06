package web

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/system"
)

func Download(
	locator string,
	output string,
) {
	r := BasicGet(locator)
	defer errors.LogClose(r.Body)
	errors.PanicStatus(r)
	f := system.Create(output)
	defer errors.LogClose(f)
	system.Copy(r.Body, f)
}
