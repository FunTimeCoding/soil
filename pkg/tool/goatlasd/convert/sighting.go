package convert

import (
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/generated/server"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/model/sighting"
)

func Sighting(v *sighting.Sighting) *server.Sighting {
	return &server.Sighting{
		Source:          v.Source,
		HardwareAddress: v.HardwareAddress,
		Address:         v.Address,
		Hostname:        &v.Hostname,
		Reserved:        v.Reserved,
		PlaceKind:       &v.PlaceKind,
		PlaceIdentifier: &v.PlaceIdentifier,
		PlaceName:       &v.PlaceName,
		Seen:            v.SeenAt,
	}
}
