package debian

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/strings"
	"github.com/funtimecoding/soil/pkg/system/debian/constant"
	"github.com/funtimecoding/soil/pkg/system/debian/image"
	"github.com/funtimecoding/soil/pkg/web"
	"github.com/funtimecoding/soil/pkg/web/locator"
)

func CheckLatestImage() {
	c := web.Client()

	for _, i := range strings.DeleteDuplicates(
		image.FindNames(
			web.GetString(
				c,
				locator.New(constant.Image).Path(
					"/cdimage/release/current/arm64/iso-cd",
				).Trail().String(),
			),
		),
	) {
		fmt.Printf("Image: %s\n", i)
	}
}
