package changed

import "github.com/funtimecoding/soil/pkg/git/constant"

func NewHook(
	directory string,
	hook string,
) *Range {
	switch hook {
	case constant.HookPrePush:
		return NewPush(directory)
	case constant.HookPreCommit:
		return NewStaged()
	default:
		return NewAll()
	}
}
