package pointer

import "net/url"

func unescape(target string) string {
	result, e := url.PathUnescape(target)

	if e != nil {
		return target
	}

	return result
}
