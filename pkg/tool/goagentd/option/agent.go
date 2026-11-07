package option

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/web/option"
	"io"
)

type Agent struct {
	option.Web
	Workspace string
	Stream    func(string, io.Writer, ...string) face.Stream
}
