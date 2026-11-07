package option

import (
	"github.com/funtimecoding/soil/pkg/iterm/face"
	"github.com/funtimecoding/soil/pkg/web/option"
	"time"
)

type Iterm struct {
	option.Web
	Source      func() face.ItermClient
	KeyInterval time.Duration
}
