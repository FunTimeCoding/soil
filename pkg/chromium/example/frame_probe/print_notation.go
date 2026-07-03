package frame_probe

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/chromium/constant"
	"github.com/funtimecoding/go-library/pkg/strings"
	"github.com/funtimecoding/go-library/pkg/system/environment"
	"github.com/funtimecoding/go-library/pkg/web"
	"github.com/funtimecoding/go-library/pkg/web/locator"
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
