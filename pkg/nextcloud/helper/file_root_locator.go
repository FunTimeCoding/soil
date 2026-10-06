package helper

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/web/locator"
)

func FileRootLocator(
	host string,
	user string,
) *locator.Locator {
	return locator.New(host).
		Base(fmt.Sprintf("remote.php/dav/files/%s", user)).
		Trail()
}
