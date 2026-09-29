package link

import "github.com/funtimecoding/soil/pkg/tool/golinkaced/generated/client"

func FromDaemonSlice(
	v []client.Link,
	host string,
) []*Link {
	var result []*Link

	for _, e := range v {
		result = append(result, FromDaemon(e, host))
	}

	return result
}
