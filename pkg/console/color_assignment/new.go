package color_assignment

import "github.com/funtimecoding/soil/pkg/face"

func New(
	name string,
	f face.SprintFunction,
) *Assignment {
	return &Assignment{Name: name, Function: f}
}
