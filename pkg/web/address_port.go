package web

import (
	"github.com/funtimecoding/soil/pkg/system/environment"
	"github.com/funtimecoding/soil/pkg/web/constant"
)

func AddressPort(p int) string {
	return AddressHostPort(
		environment.Fallback(constant.BindEnvironment, constant.Loopback),
		p,
	)
}
