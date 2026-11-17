package score_scale

import "github.com/funtimecoding/soil/pkg/face"

func GreenRed[T face.ScoreColorable](v ...T) *Scale {
	return Dynamic("#00aa00", "#aa0000", v...)
}
