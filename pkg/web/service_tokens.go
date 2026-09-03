package web

import (
	"github.com/funtimecoding/soil/pkg/system/environment"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"strings"
)

func ServiceTokens() []string {
	return strings.Split(
		environment.Required(constant.ServiceTokenEnvironment),
		",",
	)
}
