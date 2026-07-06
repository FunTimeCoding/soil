package issue

import "github.com/funtimecoding/soil/pkg/face"

func (i *Issue) SetShortStatus(f face.StringAlias) {
	i.shortStatus = f
}
