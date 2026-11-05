package option

import (
	"github.com/funtimecoding/soil/pkg/iterm/face"
	"time"
)

type Iterm struct {
	Address       string
	ServiceTokens []string
	Source        func() face.ItermClient
	KeyInterval   time.Duration
}
