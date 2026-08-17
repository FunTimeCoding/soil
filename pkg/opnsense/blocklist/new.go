package blocklist

import "github.com/funtimecoding/soil/pkg/opnsense/response"

func New(v response.Blocklist) *Blocklist {
	return &Blocklist{
		Identifier:  v.Identifier,
		Enabled:     bool(v.Enabled),
		Type:        v.Type,
		Description: v.Description,
	}
}
