package option

import (
	"github.com/funtimecoding/soil/pkg/tool/gochromed/face"
	"github.com/funtimecoding/soil/pkg/web/option"
)

type Chrome struct {
	option.Web
	DownloadDirectory string
	Source            func() face.ChromiumSource
}
