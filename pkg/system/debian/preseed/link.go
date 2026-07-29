package preseed

import (
	"github.com/funtimecoding/soil/pkg/system/constant"
	"github.com/funtimecoding/soil/pkg/web/locator"
)

func Link(release string) string {
	return locator.New(constant.DebianWeb).Path(
		"/releases/%s/example-preseed.txt",
		release,
	).String()
}
