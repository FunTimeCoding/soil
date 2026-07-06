package git

import "github.com/funtimecoding/soil/pkg/git/constant"

func Tag(name string) {
	Run(constant.Tag, name)
}
