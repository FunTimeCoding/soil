package option

import (
	"github.com/funtimecoding/soil/pkg/jellyfin"
	"github.com/funtimecoding/soil/pkg/jellyfin/face"
)

func New() *Jellyfin {
	return &Jellyfin{
		Source: func() face.JellyfinClient { return jellyfin.NewEnvironment() },
	}
}
