package age_colorer

import "github.com/funtimecoding/soil/pkg/face"

func GreenRed[T face.AgeColorable](v ...T) *Colorer {
	return Dynamic("#00aa00", "#aa0000", v...)
}
