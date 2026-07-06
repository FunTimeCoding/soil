package issue

import "github.com/funtimecoding/soil/pkg/face"

func (i *Issue) ScoreColor() face.SprintFunction {
	return i.scoreColor
}
