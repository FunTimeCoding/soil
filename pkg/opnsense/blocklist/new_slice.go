package blocklist

import "github.com/funtimecoding/soil/pkg/opnsense/response"

func NewSlice(v []response.Blocklist) []*Blocklist {
	var result []*Blocklist

	for _, e := range v {
		result = append(result, New(e))
	}

	return result
}
