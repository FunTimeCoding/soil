package system

import (
	"github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/strings/slice"
	"net"
)

func Lookup(address string) []string {
	result, e := net.LookupAddr(MustCleanAddress(address))

	if e != nil {
		return []string{}
	}

	return slice.TrimSuffix(result, constant.Dot)
}
