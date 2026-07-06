package convert

import (
	"github.com/funtimecoding/soil/pkg/netbox/prefix"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func Prefix(p *prefix.Prefix) *server.Prefix {
	result := &server.Prefix{
		Identifier: p.Identifier,
		Prefix:     p.Name,
	}

	if p.Description != "" {
		result.Description = &p.Description
	}

	if scope, okay := p.Raw.GetScope().(map[string]any); okay {
		if name, found := scope["name"].(string); found && name != "" {
			result.Site = &name
		}
	}

	return result
}
