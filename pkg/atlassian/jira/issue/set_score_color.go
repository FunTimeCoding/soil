package issue

import "github.com/funtimecoding/soil/pkg/face"

func (i *Issue) SetScoreColor(f face.SprintFunction) {
	i.scoreColor = f
}
