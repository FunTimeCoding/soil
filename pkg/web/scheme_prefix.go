package web

import "github.com/funtimecoding/soil/pkg/web/constant"

func SchemePrefix(secure bool) string {
	if secure {
		return constant.SecurePrefix
	}

	return constant.InsecurePrefix
}
