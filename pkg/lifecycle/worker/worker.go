package worker

import (
	"github.com/funtimecoding/soil/pkg/face"
	"sync"
)

type Worker struct {
	Face face.Worker
	once sync.Once
}
