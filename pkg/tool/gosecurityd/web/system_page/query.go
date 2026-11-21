package system_page

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"net/url"
)

func Query(
	path string,
	v url.Values,
) string {
	if len(v) == 0 {
		return path
	}

	return join.Empty(path, "?", v.Encode())
}
