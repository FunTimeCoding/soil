package tag

import "github.com/funtimecoding/soil/pkg/tool/golinkaced/generated/client"

func FromDaemonSlice(
	v []client.Tag,
	host string,
) []*Tag {
	var result []*Tag

	for _, e := range v {
		result = append(result, FromDaemon(e, host))
	}

	return result
}
