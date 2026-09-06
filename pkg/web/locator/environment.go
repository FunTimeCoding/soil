package locator

import (
	"github.com/funtimecoding/soil/pkg/system/environment"
	"github.com/funtimecoding/soil/pkg/web/constant"
)

func Environment(
	hostEnvironment string,
	portEnvironment string,
	insecureEnvironment string,
) *Locator {
	result := New(environment.Fallback(hostEnvironment, constant.Localhost))

	if p := environment.FallbackInteger(portEnvironment, 0); p != 0 {
		result.Port(p)
	}

	if environment.Exists(insecureEnvironment) {
		result.Insecure()
	}

	return result
}
