package option

import (
	"github.com/funtimecoding/soil/pkg/firefox/face"
	"github.com/funtimecoding/soil/pkg/web/option"
)

type Firefox struct {
	option.Web
	DownloadDirectory string
	Bridge            option.Web
	Client            func() face.FirefoxClient
}
