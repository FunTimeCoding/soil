package option

import (
	"github.com/funtimecoding/soil/pkg/jellyfin"
	"github.com/funtimecoding/soil/pkg/tool/gojellyfind/face"
)

func New() *Jellyfin {
	return &Jellyfin{
		Source: func() face.JellyfinSource { return jellyfin.NewEnvironment() },
	}
}
