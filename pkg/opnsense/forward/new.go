package forward

import "github.com/funtimecoding/soil/pkg/opnsense/response"

func New(v response.Forward) *Forward {
	return &Forward{
		Identifier:  v.Identifier,
		Enabled:     bool(v.Enabled),
		Type:        v.Type,
		Domain:      v.Domain,
		Server:      v.Server,
		Port:        v.Port,
		Description: v.Description,
	}
}
