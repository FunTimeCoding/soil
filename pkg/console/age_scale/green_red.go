package age_scale

import "github.com/funtimecoding/soil/pkg/face"

func GreenRed[T face.AgeColorable](v ...T) *Scale {
	return Dynamic("#00aa00", "#aa0000", v...)
}
