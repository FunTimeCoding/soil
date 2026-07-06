package system

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/strings/separator"
	"github.com/funtimecoding/soil/pkg/strings/slice"
	"net"
)

func MustLookup(address string) []string {
	result, e := net.LookupAddr(MustCleanAddress(address))
	errors.PanicOnError(e)

	return slice.TrimSuffix(result, separator.Dot)
}
