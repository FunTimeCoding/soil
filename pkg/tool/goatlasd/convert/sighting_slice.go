package convert

import (
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/generated/server"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/model/sighting"
)

func SightingSlice(v []*sighting.Sighting) []server.Sighting {
	result := []server.Sighting{}

	for _, e := range v {
		result = append(result, *Sighting(e))
	}

	return result
}
