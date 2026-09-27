package bookmark

import "github.com/netbox-community/go-netbox/v4"

func NewSlice(v []netbox.Bookmark) []*Bookmark {
	var result []*Bookmark

	for _, e := range v {
		result = append(result, New(&e))
	}

	return result
}
