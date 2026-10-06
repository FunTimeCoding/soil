package connection

import (
	"errors"
	"net"
	"net/url"
)

func locate(e error) (string, string) {
	var request *url.Error

	if errors.As(e, &request) {
		if u, f := url.Parse(request.URL); f == nil {
			return u.Host, u.Path
		}
	}

	var operation *net.OpError

	if errors.As(e, &operation) && operation.Addr != nil {
		return operation.Addr.String(), ""
	}

	return "", ""
}
