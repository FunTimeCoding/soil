package face

import "github.com/funtimecoding/soil/pkg/face"

type EmbedCounter interface {
	face.TokenCounter
	Close() error
}
