package service

import (
	"github.com/funtimecoding/soil/pkg/tool/goitermd/face"
	"time"
)

type Service struct {
	source      face.ItermSource
	keyInterval time.Duration
}
