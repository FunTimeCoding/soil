package list

import "github.com/funtimecoding/soil/pkg/tool/golinkaced/generated/client"

func FromDaemonSlice(
	v []client.List,
	host string,
) []*List {
	var result []*List

	for _, e := range v {
		result = append(result, FromDaemon(e, host))
	}

	return result
}
