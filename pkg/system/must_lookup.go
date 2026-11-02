package system

import (
	"github.com/funtimecoding/soil/pkg/errors"
	stringsConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/strings/slice"
	"net"
)

func MustLookup(address string) []string {
	result, e := net.LookupAddr(MustCleanAddress(address))
	errors.PanicOnError(e)

	return slice.TrimSuffix(result, stringsConstant.Dot)
}
