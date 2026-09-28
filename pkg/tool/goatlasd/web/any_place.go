package web

import "github.com/funtimecoding/soil/pkg/tool/goatlasd/model/sighting"

func AnyPlace(v []*sighting.Sighting) bool {
	for _, e := range v {
		if e.PlaceName != "" {
			return true
		}
	}

	return false
}
