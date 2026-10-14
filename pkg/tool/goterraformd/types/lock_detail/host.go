package lock_detail

import "strings"

func (d *Detail) Host() string {
	_, result, okay := strings.Cut(d.Who, "@")

	if !okay {
		return ""
	}

	return result
}
