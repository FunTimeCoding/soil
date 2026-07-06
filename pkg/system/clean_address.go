package system

import (
	"github.com/funtimecoding/soil/pkg/strings/split/key_value"
	"net"
	"strings"
)

func CleanAddress(s string) (string, error) {
	if strings.ContainsRune(s, ':') {
		host, _, e := net.SplitHostPort(s)

		if e != nil {
			return "", e
		}

		s = host
	}

	if strings.ContainsRune(s, '/') {
		host, _ := key_value.Slash(s)
		s = host
	}

	return s, nil
}
