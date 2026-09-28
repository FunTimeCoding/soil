package gazetteer

import "github.com/funtimecoding/soil/pkg/tool/goatlasd/place"

type Gazetteer struct {
	place    map[string]*place.Place
	address  map[string]*place.Place
	hardware map[string]*place.Place
}
