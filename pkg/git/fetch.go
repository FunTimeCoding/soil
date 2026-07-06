package git

import "github.com/funtimecoding/soil/pkg/git/constant"

func Fetch() {
	Run(constant.Fetch, constant.Prune, constant.PruneTags)
}
