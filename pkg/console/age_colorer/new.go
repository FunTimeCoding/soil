package age_colorer

import (
	"github.com/funtimecoding/soil/pkg/console/color_assignment"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/math/range_mapping"
)

func New[T face.AgeColorable](
	a []*color_assignment.Assignment,
	c ...T,
) *Colorer {
	var largest float64

	for _, e := range c {
		if value := e.Age().Hours(); value > largest {
			largest = value
		}
	}

	var values []float64

	for _, e := range c {
		values = append(values, e.Age().Hours()/largest)
	}

	var names []string
	assignments := make(map[string]face.SprintFunction)

	for _, e := range a {
		names = append(names, e.Name)
		assignments[e.Name] = e.Function
	}

	return &Colorer{
		assignments: assignments,
		mapping:     range_mapping.Generate(values, names),
		largest:     largest,
	}
}
