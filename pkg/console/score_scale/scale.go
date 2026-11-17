package score_scale

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/math/range_mapping"
)

type Scale struct {
	assignments map[string]face.SprintFunction
	mapping     []*range_mapping.Mapping
	largest     float64
}
