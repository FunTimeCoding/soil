package changed

import "github.com/funtimecoding/soil/pkg/git/constant"

func NewPush(directory string) *Range {
	base := Upstream(directory)

	if base == "" {
		return NewAll()
	}

	return New(base, constant.HeadReference)
}
