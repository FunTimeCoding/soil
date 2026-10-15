package configuration_context

import "github.com/funtimecoding/soil/pkg/netbox/configuration_record"

func New(v *configuration_record.Record) *Context {
	return &Context{
		Identifier:  v.Identifier,
		Name:        v.Name,
		Weight:      v.Weight,
		Description: v.Description,
		Active:      v.Active,
		Payload:     v.Payload,
		Roles:       names(v.Roles),
		Clusters:    names(v.Clusters),
		Platforms:   names(v.Platforms),
		Sites:       names(v.Sites),
		Tenants:     names(v.Tenants),
		Tags:        v.Tags,
		Raw:         v,
	}
}
