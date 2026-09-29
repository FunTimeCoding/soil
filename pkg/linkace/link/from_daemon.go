package link

import "github.com/funtimecoding/soil/pkg/tool/golinkaced/generated/client"

func FromDaemon(
	v client.Link,
	host string,
) *Link {
	return &Link{
		Identifier: int(v.Identifier),
		Title:      v.Name,
		Link:       v.Link,
		Host:       host,
	}
}
