package list

import "github.com/funtimecoding/soil/pkg/linkace/response"

func New(
	v response.List,
	host string,
) *List {
	return &List{
		Identifier:  v.Identifier,
		Name:        v.Name,
		Description: v.Description,
		Host:        host,
		Visibility:  v.Visibility,
	}
}
