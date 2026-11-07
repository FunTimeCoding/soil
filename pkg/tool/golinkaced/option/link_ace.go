package option

import (
	"github.com/funtimecoding/soil/pkg/linkace/face"
	"github.com/funtimecoding/soil/pkg/web/option"
)

type LinkAce struct {
	option.Web
	Source func() face.LinkAceClient
}
