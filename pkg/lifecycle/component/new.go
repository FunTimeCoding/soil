package component

import "github.com/funtimecoding/soil/pkg/face"

func New(w face.Worker) *Component {
	return &Component{Worker: w}
}
