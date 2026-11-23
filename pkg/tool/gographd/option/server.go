package option

import (
	"github.com/funtimecoding/soil/pkg/chat/face"
	"github.com/funtimecoding/soil/pkg/web/option"
)

type Server struct {
	option.Web
	Source func() face.TeamsClient
}
