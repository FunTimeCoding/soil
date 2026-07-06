package face

import "github.com/funtimecoding/soil/pkg/sublime/view"

type SublimeSource interface {
	Views() ([]*view.View, error)
	View(identifier int) (*view.View, error)
	CreateView(
		title string,
		content string,
		syntax string,
	) (*view.View, error)
	EditView(
		identifier int,
		old string,
		new string,
		all bool,
	) (*view.View, error)
	OpenFile(path string) (*view.View, error)
	SaveView(
		identifier int,
		path string,
	) error
	CloseView(identifier int) error
	MustViews() []*view.View
	MustView(identifier int) *view.View
	MustCreateView(
		title string,
		content string,
		syntax string,
	) *view.View
	MustEditView(
		identifier int,
		old string,
		new string,
		all bool,
	) *view.View
	MustOpenFile(path string) *view.View
}
