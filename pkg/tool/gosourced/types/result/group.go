package result

import "github.com/funtimecoding/soil/pkg/tool/gosourced/types/result/location"

type Group struct {
	Shape     string               `json:"shape"`
	Exemplar  string               `json:"exemplar"`
	Locations []*location.Location `json:"locations"`
}
