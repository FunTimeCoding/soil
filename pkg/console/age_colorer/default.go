package age_colorer

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/console/color_assignment"
	"github.com/funtimecoding/soil/pkg/face"
)

func Default[T face.AgeColorable](v ...T) *Colorer {
	return New(
		[]*color_assignment.Assignment{
			color_assignment.New(console.GreenColor, console.Green),
			color_assignment.New(console.YellowColor, console.Yellow),
			color_assignment.New(console.RedColor, console.Red),
		},
		v...,
	)
}
