package option

import "github.com/funtimecoding/soil/pkg/directory/face"

func New() *Option {
	return &Option{Directory: func() face.Client { return nil }}
}
