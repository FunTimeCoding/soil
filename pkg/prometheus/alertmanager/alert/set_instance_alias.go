package alert

import "github.com/funtimecoding/soil/pkg/face"

func (a *Alert) SetInstanceAlias(f face.StringAlias) {
	a.instance = f
}
