package web

import (
	"github.com/funtimecoding/soil/pkg/integers"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/strings/separator"
)

func AddressHostPort(
	host string,
	port int,
) string {
	return join.Empty(host, separator.Colon, integers.ToString(port))
}
