package option

import "github.com/funtimecoding/soil/pkg/tool/gochromed/face"

type Chrome struct {
	Address           string
	ServiceTokens     []string
	DownloadDirectory string
	Source            func() face.ChromiumSource
}
