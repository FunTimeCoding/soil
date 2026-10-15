package configuration_context

import "github.com/funtimecoding/soil/pkg/netbox/configuration_record"

type Context struct {
	Identifier  int32
	Name        string
	Weight      int32
	Description string
	Active      bool
	Payload     map[string]any
	Roles       []string
	Clusters    []string
	Platforms   []string
	Sites       []string
	Tenants     []string
	Tags        []string
	Raw         *configuration_record.Record
}
