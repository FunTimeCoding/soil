package instance

import (
	bandConstant "github.com/funtimecoding/soil/pkg/band/constant"
	"github.com/funtimecoding/soil/pkg/web/locator"
)

func (i *Instance) Locator() string {
	return locator.New(i.Host).Port(bandConstant.SecurePort).String()
}
