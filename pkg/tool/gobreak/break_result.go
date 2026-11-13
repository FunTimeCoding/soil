package gobreak

import (
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/tool/gobreak/break_file"
	"github.com/funtimecoding/soil/pkg/tool/gobreak/break_result"
	"github.com/funtimecoding/soil/pkg/tool/gobreak/constant"
	"github.com/funtimecoding/soil/pkg/tool/gobreak/overlay"
	"golang.org/x/tools/cover"
)

func breakResult(
	root string,
	modulePath string,
	f *break_file.File,
	k *break_file.Break,
	directory string,
	profiles []*cover.Profile,
) *break_result.Result {
	system.MakeDirectory(directory)
	result := break_result.New(k.Name, k.Expect, constant.ResultUnmatched)
	o := overlay.New(directory)
	spans, failure := apply(root, modulePath, k, o)

	if failure != "" {
		result.Detail = failure

		return result
	}

	judge(
		result,
		test(root, f, constant.OverlayFlag, o.Write()),
		spans,
		profiles,
	)
	result.Detail = o.Restore(root, result.Detail)

	return result
}
