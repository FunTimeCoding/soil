package system

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"os"
	"strings"
)

func Move(
	from string,
	to string,
) {
	if e := os.Rename(from, to); e != nil {
		// Example: "rename godownload /usr/local/bin/godownload: invalid cross-device link"
		if strings.HasSuffix(e.Error(), "invalid cross-device link") {
			MoveCopy(from, to)
		} else {
			errors.PanicOnError(e)
		}
	}
}
