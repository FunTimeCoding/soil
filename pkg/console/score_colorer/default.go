package score_colorer

import (
	"github.com/funtimecoding/soil/pkg/console/color_assignment"
	"github.com/funtimecoding/soil/pkg/console/constant"
	"github.com/funtimecoding/soil/pkg/face"
)

func Default[T face.ScoreColorable](v ...T) *Colorer {
	return New(
		[]*color_assignment.Assignment{
			color_assignment.New(constant.GreenColor, constant.Green),
			color_assignment.New(constant.YellowColor, constant.Yellow),
			color_assignment.New(constant.RedColor, constant.Red),
		},
		v...,
	)
}
