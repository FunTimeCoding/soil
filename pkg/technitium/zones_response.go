package technitium

import "github.com/funtimecoding/soil/pkg/technitium/zone"

type zonesResponse struct {
	Zones []*zone.Zone `json:"zones"`
}
