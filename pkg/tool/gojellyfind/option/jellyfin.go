package option

import (
	"github.com/funtimecoding/soil/pkg/jellyfin/face"
	"github.com/funtimecoding/soil/pkg/web/option"
)

type Jellyfin struct {
	option.Web
	Source func() face.JellyfinClient
}
