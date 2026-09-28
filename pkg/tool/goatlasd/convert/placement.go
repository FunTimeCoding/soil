package convert

import (
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/generated/server"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/model/placement"
)

func Placement(v *placement.Placement) *server.Placement {
	return &server.Placement{
		Source:          v.Source,
		Kind:            v.Kind,
		Scope:           &v.Scope,
		Name:            v.Name,
		Package:         &v.Package,
		Version:         &v.Version,
		PlaceKind:       &v.PlaceKind,
		PlaceIdentifier: &v.PlaceIdentifier,
		PlaceName:       &v.PlaceName,
		Seen:            v.SeenAt,
	}
}
