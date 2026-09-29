package tag

import "github.com/funtimecoding/soil/pkg/linkace/response"

func New(
	v response.Tag,
	host string,
) *Tag {
	return &Tag{
		Identifier: v.Identifier,
		Name:       v.Name,
		Host:       host,
		Visibility: v.Visibility,
	}
}
