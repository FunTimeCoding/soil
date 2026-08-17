package convert

import (
	"github.com/funtimecoding/soil/pkg/opnsense/blocklist"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/generated/server"
)

func Blocklist(v *blocklist.Blocklist) *server.Blocklist {
	return &server.Blocklist{
		Identifier:  v.Identifier,
		Enabled:     v.Enabled,
		Type:        v.Type,
		Description: v.Description,
	}
}
