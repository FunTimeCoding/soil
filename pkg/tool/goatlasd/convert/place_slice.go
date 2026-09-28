package convert

import (
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/generated/server"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/store/result"
)

func PlaceSlice(v []*result.Place) []server.Place {
	result := []server.Place{}

	for _, e := range v {
		result = append(result, *Place(e))
	}

	return result
}
