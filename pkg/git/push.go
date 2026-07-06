package git

import "github.com/funtimecoding/soil/pkg/git/constant"

func Push() {
	Run(constant.Push, constant.OriginRemote, constant.Tags)
}
