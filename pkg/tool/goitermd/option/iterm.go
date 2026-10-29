package option

import (
	"github.com/funtimecoding/soil/pkg/tool/goitermd/face"
	"time"
)

type Iterm struct {
	Address       string
	ServiceTokens []string
	Source        func() face.ItermSource
	KeyInterval   time.Duration
}
