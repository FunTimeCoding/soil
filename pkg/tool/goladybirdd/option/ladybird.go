package option

import (
	"github.com/funtimecoding/soil/pkg/browser/face"
	"github.com/funtimecoding/soil/pkg/browser/ladybird/types/message"
	"github.com/funtimecoding/soil/pkg/web/option"
)

type Ladybird struct {
	option.Web
	DownloadDirectory string
	Profile           string
	Headless          bool
	Client            func(events func(*message.Reply)) face.LadybirdClient
}
