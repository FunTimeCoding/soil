package option

import "github.com/funtimecoding/soil/pkg/tool/gojellyfind/face"

type Jellyfin struct {
	Address       string
	ServiceTokens []string
	Source        func() face.JellyfinSource
}
