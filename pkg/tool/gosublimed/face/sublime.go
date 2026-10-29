package face

import (
	"github.com/funtimecoding/soil/pkg/sublime/view"
	"github.com/funtimecoding/soil/pkg/tool/gosublimed/types/edit"
)

type Sublime interface {
	Views() ([]*view.View, error)
	View(identifier int) (*view.View, error)
	CreateView(
		title string,
		content string,
		syntax string,
	) (*view.View, error)
	EditView(e *edit.Edit) (*view.View, error)
	OpenFile(path string) (*view.View, error)
	SaveView(
		identifier int,
		path string,
	) error
	CloseView(identifier int) error
}
