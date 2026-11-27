package chromium

import (
	"github.com/funtimecoding/soil/pkg/browser/constant"
	"github.com/funtimecoding/soil/pkg/strings"
	"github.com/funtimecoding/soil/pkg/system/environment"
)

func ListeningEnvironment() bool {
	return Listening(
		environment.Required(constant.ChromiumHostEnvironment),
		strings.MustToInteger(
			environment.Required(constant.ChromiumPortEnvironment),
		),
	)
}
