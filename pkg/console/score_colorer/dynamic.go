package score_colorer

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/console/color_assignment"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/integers"
)

func Dynamic[T face.ScoreColorable](
	startHex string,
	endHex string,
	v ...T,
) *Colorer {
	count := len(v)
	c := console.Gradient(startHex, endHex, count)
	assignments := make([]*color_assignment.Assignment, count)

	for i := 0; i < count; i++ {
		assignments[i] = color_assignment.New(integers.ToString(i), c[i])
	}

	return New(assignments, v...)
}
