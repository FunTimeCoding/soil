package response

import "github.com/funtimecoding/soil/pkg/technitium/zone"

type Zones struct {
	Zones []*zone.Zone `json:"zones"`
}
