package gofix

import (
	"github.com/funtimecoding/soil/pkg/lint/face"
	"github.com/funtimecoding/soil/pkg/source/index"
	"github.com/funtimecoding/soil/pkg/tool/gofix/option"
)

func indexHandle(o *option.Fix) *index.Workspace {
	if o.Handle != nil {
		return o.Handle
	}

	return index.New(o.IndexDirectory(), o.Root, face.Kind())
}
