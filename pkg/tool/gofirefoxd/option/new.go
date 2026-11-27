package option

import (
	"github.com/funtimecoding/soil/pkg/browser/face"
	"github.com/funtimecoding/soil/pkg/browser/firefox"
)

func New() *Firefox {
	return &Firefox{Client: func() face.FirefoxClient { return firefox.New() }}
}
