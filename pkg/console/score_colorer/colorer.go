package score_colorer

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/math/range_mapping"
)

type Colorer struct {
	assignments map[string]face.SprintFunction
	mapping     []*range_mapping.Mapping
	largest     float64
}
