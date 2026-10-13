package frame_probe

import (
	"github.com/funtimecoding/soil/pkg/chromium/constant"
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/strings"
	"github.com/funtimecoding/soil/pkg/system/environment"
	"github.com/funtimecoding/soil/pkg/web"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"github.com/funtimecoding/soil/pkg/web/stall_client"
)

func printNotation() {
	console.Line("=== raw /json ===")
	p := stall_client.New().Insecure().Build()
	defer p.Close()
	console.Line(
		web.GetString(
			p.Client(),
			locator.New(environment.Required(constant.HostEnvironment)).Port(
				strings.MustToInteger(
					environment.Required(constant.PortEnvironment),
				),
			).Path(constant.NotationPath).Insecure().String(),
		),
	)
}
