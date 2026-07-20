package board

import "github.com/funtimecoding/soil/pkg/tool/godashboardd/constant"

func Port(t Target) int {
	if t.Port != 0 {
		return t.Port
	}

	if t.Secure {
		return constant.SecurePort
	}

	return constant.InsecurePort
}
