package face

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/sublime/view"
)

type SublimeClient interface {
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
	Address() string
	face.Reachable
	Close()
}
