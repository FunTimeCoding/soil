package web

import (
	"github.com/funtimecoding/soil/pkg/integers"
	stringsConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
)

func AddressHostPort(
	host string,
	port int,
) string {
	return join.Empty(host, stringsConstant.Colon, integers.ToString(port))
}
