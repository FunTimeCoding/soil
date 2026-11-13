package entry

import (
	"github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/extended"
	"maragu.dev/gomponents"
)

func (e *Entry) extendedAttributes() []gomponents.Node {
	var result []gomponents.Node

	if e.target != "" {
		request := extended.Post(e.path)

		if e.query {
			request = extended.Get(e.path)
		}

		result = append(
			result,
			request,
			extended.Target(e.target),
			extended.Swap(e.swap),
		)

		if e.live {
			result = append(
				result,
				extended.Trigger(constant.TriggerForm),
				extended.Sync(constant.SyncReplace),
				extended.Push("true"),
			)
		}
	} else if e.previewPath != "" {
		result = append(
			result,
			extended.Post(e.previewPath),
			extended.Trigger(constant.EntryPreviewDelay),
			extended.Target(e.previewTarget),
			extended.Swap(constant.SwapOuter),
		)
	}

	if e.reset {
		result = append(
			result,
			extended.AfterRequest(constant.EntryResetScript),
		)
	}

	return result
}
