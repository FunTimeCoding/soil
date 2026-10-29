package option

import "github.com/funtimecoding/soil/pkg/firefox/face"

type Firefox struct {
	Address           string
	DownloadDirectory string
	ServiceTokens     []string
	BridgePort        int
	Client            func() face.Client
}
