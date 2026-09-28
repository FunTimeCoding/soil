package convert

import (
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/generated/server"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/model/placement"
)

func PlacementSlice(v []*placement.Placement) []server.Placement {
	result := []server.Placement{}

	for _, e := range v {
		result = append(result, *Placement(e))
	}

	return result
}
