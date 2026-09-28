package convert

import (
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/generated/server"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/store/result"
)

func Place(v *result.Place) *server.Place {
	return &server.Place{
		Name:       v.Name,
		Kind:       v.Kind,
		Identifier: &v.Identifier,
		Count:      v.Count,
	}
}
