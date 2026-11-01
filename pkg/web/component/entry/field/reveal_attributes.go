package field

import (
	"github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/extended"
	"maragu.dev/gomponents"
)

func (f *Field) revealAttributes() gomponents.Node {
	if f.revealPath == "" {
		return gomponents.Group(nil)
	}

	return gomponents.Group(
		[]gomponents.Node{
			extended.Get(f.revealPath),
			extended.Trigger(constant.EntryRevealTrigger),
			extended.Target(f.revealTarget),
			extended.Swap(constant.SwapInner),
			extended.Include(constant.EntryRevealInclude),
		},
	)
}
