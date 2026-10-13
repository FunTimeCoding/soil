package install

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/strings"
	"github.com/funtimecoding/soil/pkg/system/constant"
	"github.com/funtimecoding/soil/pkg/system/linux/debian/install/image"
	"github.com/funtimecoding/soil/pkg/web"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"github.com/funtimecoding/soil/pkg/web/stall_client"
)

func CheckLatestImage() {
	p := stall_client.New().Build()
	defer p.Close()
	c := p.Client()

	for _, i := range strings.DeleteDuplicates(
		image.FindNames(
			web.GetString(
				c,
				locator.New(constant.DebianImage).Path(
					"/cdimage/release/current/arm64/iso-cd",
				).Trail().String(),
			),
		),
	) {
		fmt.Printf("Image: %s\n", i)
	}
}
