package web

import (
	"github.com/funtimecoding/soil/pkg/integers"
	"github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
)

func AddressHostPort(
	host string,
	port int,
) string {
	return join.Empty(host, constant.Colon, integers.ToString(port))
}
