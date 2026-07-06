package frame_probe

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/chromium/constant"
	"github.com/funtimecoding/soil/pkg/strings"
	"github.com/funtimecoding/soil/pkg/system/environment"
	"github.com/funtimecoding/soil/pkg/web"
	"github.com/funtimecoding/soil/pkg/web/locator"
)

func printNotation() {
	fmt.Println("=== raw /json ===")
	fmt.Println(
		web.GetString(
			web.InsecureClient(),
			locator.New(
				environment.Required(constant.HostEnvironment),
			).Port(
				strings.MustToInteger(
					environment.Required(constant.PortEnvironment),
				),
			).Path(constant.NotationPath).Insecure().String(),
		),
	)
}
