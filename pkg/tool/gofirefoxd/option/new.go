package option

import (
	"github.com/funtimecoding/soil/pkg/firefox"
	"github.com/funtimecoding/soil/pkg/firefox/face"
)

func New() *Firefox {
	return &Firefox{Client: func() face.Client { return firefox.New() }}
}
