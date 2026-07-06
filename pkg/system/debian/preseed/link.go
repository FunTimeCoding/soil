package preseed

import (
	"github.com/funtimecoding/soil/pkg/system/debian/constant"
	"github.com/funtimecoding/soil/pkg/web/locator"
)

func Link(release string) string {
	return locator.New(constant.Web).Path(
		"/releases/%s/example-preseed.txt",
		release,
	).String()
}
