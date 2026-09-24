package gohook

import (
	"github.com/funtimecoding/soil/pkg/git/changed"
	"github.com/funtimecoding/soil/pkg/git/constant"
	"github.com/funtimecoding/soil/pkg/tool/gohook/option"
)

func ResolveRange(
	root string,
	o *option.Run,
) *changed.Range {
	switch {
	case o.Base != "":
		return changed.New(o.Base, o.Head)
	case o.Hook == constant.HookPrePush && o.Input != "":
		return changed.NewPushRefs(root, o.Input)
	default:
		return changed.NewHook(root, o.Hook)
	}
}
