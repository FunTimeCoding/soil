package convert

import (
	"github.com/funtimecoding/soil/pkg/opnsense/alias"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/generated/server"
)

func Alias(v *alias.Alias) *server.Alias {
	return &server.Alias{
		Identifier:  v.Identifier,
		Enabled:     v.Enabled,
		Name:        v.Name,
		Type:        v.Type,
		Content:     v.Content,
		Description: v.Description,
	}
}
