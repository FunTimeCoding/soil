package convert

import (
	"github.com/funtimecoding/soil/pkg/opnsense/forward"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/generated/server"
)

func Forward(v *forward.Forward) *server.Forward {
	return &server.Forward{
		Identifier:  v.Identifier,
		Enabled:     v.Enabled,
		Type:        v.Type,
		Domain:      v.Domain,
		Server:      v.Server,
		Port:        v.Port,
		Description: v.Description,
	}
}
