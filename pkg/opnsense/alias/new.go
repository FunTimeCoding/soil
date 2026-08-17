package alias

import "github.com/funtimecoding/soil/pkg/opnsense/response"

func New(v response.Alias) *Alias {
	return &Alias{
		Identifier:  v.Identifier,
		Enabled:     bool(v.Enabled),
		Name:        v.Name,
		Type:        v.Type,
		Content:     v.Content,
		Description: v.Description,
	}
}
