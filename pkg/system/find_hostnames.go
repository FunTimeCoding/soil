package system

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/constant"
	"net"
	"strings"
)

func FindHostnames(host string) ([]string, error) {
	var result []string
	clean, cleanFail := CleanAddress(host)

	if cleanFail != nil {
		return result, cleanFail
	}

	host = clean

	if net.ParseIP(host) != nil {
		names, lookUpFail := net.LookupAddr(host)

		if lookUpFail != nil {
			if strings.HasSuffix(
				lookUpFail.Error(),
				"nodename nor servname provided, or not known",
			) {
				return result, nil
			}

			return result, fmt.Errorf("lookup: %s", lookUpFail)
		}

		result = names

		for i := range result {
			result[i] = strings.TrimSuffix(
				result[i],
				constant.CurrentDirectory,
			)
		}
	} else {
		result = []string{host}
	}

	return result, nil
}
