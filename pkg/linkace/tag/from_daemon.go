package tag

import "github.com/funtimecoding/soil/pkg/tool/golinkaced/generated/client"

func FromDaemon(
	v client.Tag,
	host string,
) *Tag {
	return &Tag{
		Identifier: int(v.Identifier),
		Name:       v.Name,
		Host:       host,
	}
}
