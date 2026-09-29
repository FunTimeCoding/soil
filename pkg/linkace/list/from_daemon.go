package list

import "github.com/funtimecoding/soil/pkg/tool/golinkaced/generated/client"

func FromDaemon(
	v client.List,
	host string,
) *List {
	return &List{
		Identifier: int(v.Identifier),
		Name:       v.Name,
		Host:       host,
	}
}
