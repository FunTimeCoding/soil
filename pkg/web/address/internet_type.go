package address

import (
	"github.com/funtimecoding/soil/pkg/web/constant"
	"net"
)

func InternetType(s string) string {
	i := net.ParseIP(s)

	if i == nil {
		return constant.AddressNoneType
	}

	if i.To4() != nil {
		return constant.AddressFourType
	}

	return constant.AddressSixType
}
