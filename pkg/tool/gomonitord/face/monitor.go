package face

import "github.com/funtimecoding/soil/pkg/face"

type Monitor interface {
	Claims
	face.EventNotifier
}
