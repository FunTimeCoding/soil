package face

import "github.com/funtimecoding/soil/pkg/face"

type ClaimStream interface {
	ClaimSource
	face.EventNotifier
}
