package option

import "github.com/funtimecoding/soil/pkg/jellyfin/face"

type Jellyfin struct {
	Address       string
	ServiceTokens []string
	Source        func() face.JellyfinClient
}
