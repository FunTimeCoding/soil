package convert

import "net"

func address(v []*net.IPNet) []string {
	result := make([]string, 0, len(v))

	for _, a := range v {
		result = append(result, a.String())
	}

	return result
}
