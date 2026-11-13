package gofix

import (
	"github.com/funtimecoding/soil/pkg/tool/gofix/option"
	"github.com/funtimecoding/soil/pkg/tool/gofix/workspace"
)

func workspaceFor(o *option.Fix) *workspace.Workspace {
	if o.Layer != nil {
		return workspace.Over(o.Layer, o.Diff)
	}

	return workspace.New(o.Diff, append([]string{o.Root}, o.Replacing...)...)
}
