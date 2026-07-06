package issue

import "github.com/funtimecoding/soil/pkg/face"

func (i *Issue) ShortStatus() face.StringAlias {
	return i.shortStatus
}
